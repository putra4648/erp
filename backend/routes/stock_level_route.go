package routes

import (
	"putra4648/erp/configs/middleware"
	"putra4648/erp/internal/stock_level/dto"
	"putra4648/erp/internal/stock_level/mapper"
	"putra4648/erp/internal/stock_level/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func RegisterStockLevelRoutes(
	app *fiber.App,
	api fiber.Router,
	stockLevelService service.StockLevelQueryService,
) {
	stock := api.Group("/stock-levels")
	{
		stock.Get("/", middleware.RequirePermission("read:stock-levels"), getAllStockLevels(stockLevelService))
		stock.Get("/:id", middleware.RequirePermission("read:stock-levels"), getStockLevelByID(stockLevelService))
		stock.Get("/:wh_id/:prod_id", middleware.RequirePermission("read:stock-levels"), getStockLevelByWarehouseAndProduct(stockLevelService))
	}
}

func getAllStockLevels(s service.StockLevelQueryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req dto.StockLevelRequest
		if err := c.QueryParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		var warehouseIDPtr *uuid.UUID
		if req.WarehouseID != "" {
			if id, err := uuid.Parse(req.WarehouseID); err == nil {
				warehouseIDPtr = &id
			}
		}

		var productIDPtr *uuid.UUID
		if req.ProductID != "" {
			if id, err := uuid.Parse(req.ProductID); err == nil {
				productIDPtr = &id
			}
		}

		res, total, err := s.GetAllStockLevels(c.Context(), warehouseIDPtr, productIDPtr, req.Search, req.Page, req.Size)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{
			"items": mapper.ToStockLevelResponses(res),
			"total": total,
			"page":  req.Page,
			"size":  req.Size,
		})
	}
}

func getStockLevelByID(s service.StockLevelQueryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
		}

		res, err := s.FindByID(c.Context(), id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Stock Level not found"})
		}

		return c.JSON(mapper.ToStockLevelResponse(res))
	}
}

func getStockLevelByWarehouseAndProduct(s service.StockLevelQueryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		whID, err := uuid.Parse(c.Params("wh_id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Warehouse ID"})
		}

		prodID, err := uuid.Parse(c.Params("prod_id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Product ID"})
		}

		res, err := s.GetByProductAndWarehouseWithPreload(c.Context(), prodID, whID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Stock Level not found"})
		}

		return c.JSON(mapper.ToStockLevelResponse(res))
	}
}
