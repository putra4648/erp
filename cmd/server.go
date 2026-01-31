package app

import (
	"putra4648/erp/configs/config"
	"putra4648/erp/configs/logger"
	"putra4648/erp/configs/middleware"
	supplierService "putra4648/erp/internal/modules/inventory/supplier/service"
	warehouseService "putra4648/erp/internal/modules/inventory/warehouse/service"
	"putra4648/erp/routes"

	"github.com/casbin/casbin/v3"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AppDependencies struct {
	dig.In
	Config                  *config.AppEnv
	DB                      *gorm.DB
	Enforcer                *casbin.Enforcer
	Verifier                *oidc.IDTokenVerifier
	ZapLogger               *zap.Logger
	WarehouseCommandService warehouseService.WarehouseCommandService
	WarehouseQueryService   warehouseService.WarehouseQueryService
	SupplierCommandService  supplierService.SupplierCommandService
	SupplierQueryService    supplierService.SupplierQueryService
}

func Server(deps AppDependencies) error {
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
