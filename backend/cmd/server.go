package app

import (
	"putra4648/erp/configs/config"
	"putra4648/erp/configs/logger"
	"putra4648/erp/configs/middleware"
	categoryService "putra4648/erp/internal/modules/category/service"
	productService "putra4648/erp/internal/modules/product/service"
	stockAdjustmentService "putra4648/erp/internal/modules/stock_adjustment/service"
	stockMovementService "putra4648/erp/internal/modules/stock_movement/service"
	supplierService "putra4648/erp/internal/modules/supplier/service"
	uomService "putra4648/erp/internal/modules/uom/service"
	warehouseService "putra4648/erp/internal/modules/warehouse/service"
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
	ProductCommandService   productService.ProductCommandService
	ProductQueryService     productService.ProductQueryService
	UOMQueryService         uomService.UOMQueryService
	UOMCommandService       uomService.UOMCommandService
	CategoryQueryService    categoryService.CategoryQueryService
	CategoryCommandService  categoryService.CategoryCommandService

	StockAdjustmentService  stockAdjustmentService.StockAdjustmentService
	AdjustmentReasonService stockAdjustmentService.AdjustmentReasonService

	StockMovementCommandService stockMovementService.StockMovementCommandService
	StockMovementQueryService   stockMovementService.StockMovementQueryService
}

func Server(deps AppDependencies) error {
	defer deps.ZapLogger.Sync()

	// Add policies: p, role, path, action
	// This gives the 'admin' role GET access to all routes under /api/admin/*
	if hasPolicy, _ := deps.Enforcer.HasPolicy("admin", "/api/admin/*", "GET"); !hasPolicy {
		if _, err := deps.Enforcer.AddPolicy("admin", "/api/admin/*", "GET"); err != nil {
			logger.Log.Warnf("Could not add admin policy: %v", err)
		}
	}

	// Add product-specific policies for different roles
	productPolicies := []struct {
		role string
		obj  string
		act  string
	}{
		{"admin", "/api/products/*", "GET"},
		{"admin", "/api/products/*", "POST"},
		{"admin", "/api/products/*", "PUT"},
		{"admin", "/api/products/*", "DELETE"},
		{"manager", "/api/products/*", "GET"},
		{"manager", "/api/products/*", "POST"},
		{"staff", "/api/products/*", "GET"},
	}

	for _, policy := range productPolicies {
		if hasPolicy, _ := deps.Enforcer.HasPolicy(policy.role, policy.obj, policy.act); !hasPolicy {
			if _, err := deps.Enforcer.AddPolicy(policy.role, policy.obj, policy.act); err != nil {
				logger.Log.Warnf("Could not add product policy for %s: %v", policy.role, err)
			}
		}
	}

	sqlDb, err := deps.DB.DB()
	if err != nil {
		return err
	}
	sqlDb.SetMaxIdleConns(5)
	sqlDb.SetMaxOpenConns(20)

	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.GlobalErrorHandler(deps.ZapLogger),
	})

	// Recover Middleware
	app.Use(middleware.RecoverMiddleware(deps.ZapLogger))

	// Logger Middleware
	app.Use(middleware.LoggerMiddleware(deps.ZapLogger))

	// Public Route
	app.Get("/api/ping", func(c *fiber.Ctx) error { return c.JSON("pong") })

	// Protected Route (Semua user yang login)
	api := app.Group("/api")
	api.Use(middleware.AuthMiddleware(deps.Verifier))

	routes.RegisterAdminRoutes(app, api, deps.Enforcer)
	routes.RegisterUserProfile(app, api)
	routes.RegisterWarehouseRoutes(api, deps.WarehouseCommandService, deps.WarehouseQueryService)
	routes.RegisterSupplierRoutes(api, deps.SupplierCommandService, deps.SupplierQueryService)
	routes.RegisterProductRoutes(app, api, deps.ProductCommandService, deps.ProductQueryService, deps.Enforcer)
	routes.RegisterCategoryRoutes(app, api, deps.CategoryCommandService, deps.CategoryQueryService, deps.Enforcer)
	routes.RegisterUOMRoutes(app, api, deps.UOMCommandService, deps.UOMQueryService, deps.Enforcer)
	routes.RegisterStockAdjustmentRoutes(api, deps.StockAdjustmentService, deps.AdjustmentReasonService)
	routes.RegisterStockMovementRoutes(api, deps.StockMovementCommandService, deps.StockMovementQueryService)

	return app.Listen(":" + deps.Config.Port)
}
