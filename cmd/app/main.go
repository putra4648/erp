package main

import (
	"putra4648/erp/configs/auth"
	"putra4648/erp/configs/config"
	"putra4648/erp/configs/database"
	"putra4648/erp/configs/logger"
	"putra4648/erp/configs/middleware"
	"putra4648/erp/internal/modules/inventory"
	"putra4648/erp/routes"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

func main() {
	container := dig.New()

	// Register constructors
	providers := []interface{}{
		config.LoadConfig,
		logger.InitLogger,
		database.InitDatabase,
		auth.SetupCasbin,
		auth.NewOIDCProvider,
		auth.NewOIDCVerifier,
	}

	for _, p := range providers {
		if err := container.Provide(p); err != nil {
			logger.Log.Fatalf("Failed to provide dependency: %v", err)
		}
	}

	if err := inventory.Register(container); err != nil {
		logger.Log.Fatalf("Failed to register inventory module: %v", err)
	}

	container.Provide(func() AppDependencies {
		return AppDependencies{}
	})

	// Invoke the application runner
	if err := container.Invoke(run); err != nil {
		logger.Log.Fatalf("Application failed: %v", err)
	}
}

func run(deps AppDependencies) error {
	defer deps.ZapLogger.Sync()

	// Add policies: p, role, path, action
	// This gives the 'admin' role GET access to all routes under /api/admin
	if hasPolicy, _ := deps.Enforcer.HasPolicy("admin", "/api/admin/*", "GET"); !hasPolicy {
		if _, err := deps.Enforcer.AddPolicy("admin", "/api/admin/*", "GET"); err != nil {
			logger.Log.Warnf("Could not add admin policy: %v", err)
		}
	}

	sqlDb, err := deps.DB.DB()
	if err != nil {
		return err
	}
	sqlDb.SetMaxIdleConns(5)
	sqlDb.SetMaxOpenConns(20)

	app := fiber.New()

	// Public Route
	app.Get("/api/ping", func(c *fiber.Ctx) error { return c.JSON("pong") })

	// Protected Route (Semua user yang login)
	api := app.Group("/api")
	api.Use(middleware.AuthMiddleware(deps.Verifier))

	routes.RegisterAdminRoutes(app, api, deps.Enforcer)
	routes.RegisterUserProfile(app, api)
	routes.RegisterInventoryRoutes(app, api, deps.WarehouseCommandService, deps.WarehouseQueryService, deps.SupplierCommandService, deps.SupplierQueryService)

	return app.Listen(":" + deps.Config.Port)
}
