package main

import (
	"context"
	"os"
	"putra4648/erp/middleware"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)

	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	// Configuration from Environment Variables
	keycloakURL := os.Getenv("KEYCLOAK_URL")
	clientID := os.Getenv("KEYCLOAK_CLIENT_ID")
	dbDSN := os.Getenv("DB_DSN")

	// Initialize OIDC Provider
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, keycloakURL)
	if err != nil {
		logrus.Fatalf("Gagal inisialisasi OIDC provider: %v", err)
	}
	verifier := provider.Verifier(&oidc.Config{ClientID: clientID})

	app := fiber.New()

	// Public Route
	app.Get("/api/ping", func(c *fiber.Ctx) error { return c.JSON("pong") })

	// Protected Route (Semua user yang login)
	api := app.Group("/api")
	api.Use(middleware.AuthMiddleware(verifier))
	{
		api.Get("/profile", func(c *fiber.Ctx) error {
			uid := c.Locals("user_id")
			return c.JSON(fiber.Map{"user_id": uid})
		})

		// Route khusus admin
		admin := api.Group("/admin")
		admin.Use(middleware.RoleMiddleware("admin")) // Nama role yang Anda buat di Keycloak
		{
			admin.Get("/dashboard", func(c *fiber.Ctx) error {
				return c.JSON(fiber.Map{"status": "Welcome Admin!"})
			})
		}
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbDSN,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		logrus.Fatalf("Cannot connect to DB: %v", err)
	}

	logrus.Info("DB Connected ", db)
	logrus.Fatal(app.Listen(":8080"))

}
