package routes

import (
	"putra4648/erp/internal/modules/stock_movement/dto"
	"putra4648/erp/internal/modules/stock_movement/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func RegisterStockMovementRoutes(
	api fiber.Router,
	commandService service.StockMovementCommandService,
	queryService service.StockMovementQueryService,
) {
	movement := api.Group("/stock-movement")
	{
		movement.Post("/", createStockMovement(commandService))
		movement.Get("/:id", getStockMovementByID(queryService))
		movement.Get("/", getAllStockMovements(queryService))
		movement.Put("/:id", updateStockMovement(commandService))
		movement.Delete("/:id", deleteStockMovement(commandService))
	}
}

func createStockMovement(s service.StockMovementCommandService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req dto.StockMovementDTO
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		res, err := s.Create(c.Context(), &req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusCreated).JSON(res)
	}
}

func getStockMovementByID(s service.StockMovementQueryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
		}
		res, err := s.FindByID(c.Context(), id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(res)
	}
}

func getAllStockMovements(s service.StockMovementQueryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		page := c.QueryInt("page", 1)
		size := c.QueryInt("size", 10)
		movementType := c.Query("type")

		req := &dto.StockMovementRequest{
			Page: page,
			Size: size,
			Type: movementType,
		}

		res, total, err := s.FindAll(c.Context(), req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(fiber.Map{
			"items": res,
			"total": total,
			"page":  page,
			"size":  size,
		})
	}
}

func updateStockMovement(s service.StockMovementCommandService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
		}

		var req dto.StockMovementDTO
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		res, err := s.Update(c.Context(), id, &req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(res)
	}
}

func deleteStockMovement(s service.StockMovementCommandService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
		}

		if err := s.Delete(c.Context(), id); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.SendStatus(fiber.StatusNoContent)
	}
}
