package routes

import (
	"putra4648/erp/configs/middleware"
	sharedDto "putra4648/erp/internal/shared/dto"
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

		pagination := sharedDto.PaginationRequest{
			Page: c.QueryInt("page", 1),
			Size: c.QueryInt("size", 10),
		}
		var req *dto.StockLevelDto
		if err := c.QueryParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		res, err := s.FindAllStockLevels(c.Context(), &pagination, req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(res)
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
		whID := c.Params("wh_id")
		prodID := c.Params("prod_id")

		res, err := s.FindByProductAndWarehouseWithPreload(c.Context(), &dto.StockLevelDto{
			WarehouseID: &whID,
			ProductID:   &prodID,
		})
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Stock Level not found"})
		}

		return c.JSON(mapper.ToStockLevelResponse(res))
	}
}
