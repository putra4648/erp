package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"putra4648/erp/middleware"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Configuration from Environment Variables
	keycloakURL := os.Getenv("KEYCLOAK_URL")
	clientID := os.Getenv("KEYCLOAK_CLIENT_ID")
	dbDSN := os.Getenv("DB_DSN")

	// Initialize OIDC Provider
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, keycloakURL)
	if err != nil {
		log.Fatalf("Gagal inisialisasi OIDC provider: %v", err)
	}
	verifier := provider.Verifier(&oidc.Config{ClientID: clientID})

	app := gin.Default()

	// Public Route
	app.GET("/api/ping", func(c *gin.Context) { c.JSON(200, "pong") })

	// Protected Route (Semua user yang login)
	api := app.Group("/api")
	api.Use(middleware.AuthMiddleware(verifier))
	{
		api.GET("/profile", func(c *gin.Context) {
			uid := c.MustGet("user_id")
			c.JSON(200, gin.H{"user_id": uid})
		})

		// Route khusus admin
		admin := api.Group("/admin")
		admin.Use(middleware.RoleMiddleware("admin")) // Nama role yang Anda buat di Keycloak
		{
			admin.GET("/dashboard", func(c *gin.Context) {
				c.JSON(200, gin.H{"status": "Welcome Admin!"})
			})
		}
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbDSN,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("Cannot connect to DB: %v", err)
	}

	fmt.Println("DB Connected %s", db)
	app.Run()

}
