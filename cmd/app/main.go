package main

import (
	"context"
	"log"
	"putra4648/erp/configs/auth"
	"putra4648/erp/configs/config"
	"putra4648/erp/configs/logger"
	"putra4648/erp/configs/middleware"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	// initialize config
	cfg := config.LoadConfig()

	// Initialize logger
	zapLogger, err := logger.InitLogger()
	if err != nil {
		log.Fatalf("could not initialize logger: %v", err)
	}
	defer zapLogger.Sync()

	// Initialize OIDC Provider
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, cfg.KeycloakURL+"/realms/"+cfg.KeycloakRealmName)
	if err != nil {
		logger.Log.Fatalf("Failed to initialize OIDC provider: %v", err)
	}

	verifier := provider.Verifier(&oidc.Config{ClientID: cfg.KeycloakClientID})

	app := fiber.New()

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  cfg.DBDSN,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		logger.Log.Fatalf("Cannot connect to DB: %v", err)
	}

	// Initialize Casbin Enforcer
	enforcer, err := auth.SetupCasbin(db)
	if err != nil {
		logger.Log.Fatalf("Failed to setup Casbin: %v", err)
	}

	// Add policies: p, role, path, action
	// This gives the 'admin' role GET access to all routes under /api/admin
	if hasPolicy, _ := enforcer.HasPolicy("admin", "/api/admin/*", "GET"); !hasPolicy {
		if _, err := enforcer.AddPolicy("admin", "/api/admin/*", "GET"); err != nil {
			logger.Log.Warnf("Could not add admin policy: %v", err)
		}
	}

	sqlDb, err := db.DB()
	if err != nil {
		logger.Log.Fatalf("Cannot get DB: %v", err)
	}
	sqlDb.SetMaxIdleConns(5)
	sqlDb.SetMaxOpenConns(20)

	logger.Log.Info("DB Connected")

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

		// Route khusus admin, protected by Casbin
		// NOTE: The AuthMiddleware should be configured to extract the user's roles
		// from the JWT claims (e.g., from 'realm_access.roles') and place them
		// in c.Locals("roles") as a []string for the PermissionMiddleware to use.
		admin := api.Group("/admin")
		admin.Use(middleware.PermissionMiddleware(enforcer)) // Use Casbin for authorization
		{
			admin.Get("/dashboard", func(c *fiber.Ctx) error {
				return c.JSON(fiber.Map{"status": "Welcome Admin!"})
			})
		}
	}

	app.Listen(":" + cfg.Port)

}
