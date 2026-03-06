package app

import (
	"putra4648/erp/configs/auth"
	"putra4648/erp/configs/config"
	"putra4648/erp/configs/middleware"
	categoryService "putra4648/erp/internal/category/service"
	productService "putra4648/erp/internal/product/service"
	stockAdjustmentService "putra4648/erp/internal/stock_adjustment/service"
	stockLevelService "putra4648/erp/internal/stock_level/service"
	stockMovementService "putra4648/erp/internal/stock_movement/service"
	supplierService "putra4648/erp/internal/supplier/service"
	uomService "putra4648/erp/internal/uom/service"
	warehouseService "putra4648/erp/internal/warehouse/service"
	"putra4648/erp/routes"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AppDependencies struct {
	dig.In
	Config                  *config.AppEnv
	DB                      *gorm.DB
	Authenticator           *auth.Authenticator
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

	StockAdjustmentQueryService    stockAdjustmentService.StockAdjustmentQueryService
	StockAdjustmentCommandService  stockAdjustmentService.StockAdjustmentCommandService
	AdjustmentReasonQueryService   stockAdjustmentService.AdjustmentReasonQueryService
	AdjustmentReasonCommandService stockAdjustmentService.AdjustmentReasonCommandService

	StockLevelService stockLevelService.StockLevelQueryService

	StockMovementCommandService stockMovementService.StockMovementCommandService
	StockMovementQueryService   stockMovementService.StockMovementQueryService
}

func Server(deps AppDependencies) error {
	defer deps.ZapLogger.Sync()

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
	api.Use(middleware.AuthMiddleware(deps.Authenticator, deps.ZapLogger))

	routes.RegisterUserProfile(app, api)
	routes.RegisterWarehouseRoutes(api, deps.WarehouseCommandService, deps.WarehouseQueryService)
	routes.RegisterSupplierRoutes(api, deps.SupplierCommandService, deps.SupplierQueryService)
	routes.RegisterProductRoutes(app, api, deps.ProductCommandService, deps.ProductQueryService)
	routes.RegisterCategoryRoutes(app, api, deps.CategoryCommandService, deps.CategoryQueryService)
	routes.RegisterUOMRoutes(app, api, deps.UOMCommandService, deps.UOMQueryService)
	routes.RegisterStockAdjustmentRoutes(api, deps.StockAdjustmentQueryService, deps.StockAdjustmentCommandService, deps.AdjustmentReasonQueryService, deps.AdjustmentReasonCommandService)
	routes.RegisterStockMovementRoutes(api, deps.StockMovementCommandService, deps.StockMovementQueryService)
	routes.RegisterStockLevelRoutes(app, api, deps.StockLevelService)

	return app.Listen(":" + deps.Config.Port)
}
