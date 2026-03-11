package routes

import (
	"putra4648/erp/configs/middleware"
	sharedDto "putra4648/erp/internal/shared/dto"
	"putra4648/erp/internal/stock_movement/dto"
	"putra4648/erp/internal/stock_movement/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func RegisterStockMovementRoutes(
	api fiber.Router,
	commandService service.StockMovementCommandService,
	queryService service.StockMovementQueryService,
) {
	movements := api.Group("/stock-movements")
	{
		movements.Post("/", middleware.RequirePermission("create:movements"), createStockMovement(commandService))
		movements.Get("/", middleware.RequirePermission("read:movements"), getAllStockMovements(queryService))
		movements.Get("/transactions", middleware.RequirePermission("read:transactions"), getStockTransactions(queryService))
		movements.Get("/:id", middleware.RequirePermission("read:movements"), getStockMovementByID(queryService))
		movements.Put("/:id", middleware.RequirePermission("update:movements"), updateStockMovement(commandService))
		movements.Delete("/:id", middleware.RequirePermission("delete:movements"), deleteStockMovement(commandService))
		movements.Post("/:id/approve", middleware.RequirePermission("approve:movements"), approveStockMovement(commandService))
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
		var req dto.StockMovementDTO
		if err := c.QueryParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		res, err := s.FindAll(c.Context(), &sharedDto.PaginationRequest{Page: c.QueryInt("page", 1), Size: c.QueryInt("size", 10)}, &req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(res)
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

func approveStockMovement(s service.StockMovementCommandService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid movement ID"})
		}

		if err := s.Approve(c.Context(), id); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Movement approved successfully"})
	}
}

func getStockTransactions(s service.StockMovementQueryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req dto.StockTransactionDTO
		if err := c.QueryParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		res, err := s.FindTransactions(c.Context(), &sharedDto.PaginationRequest{Page: c.QueryInt("page", 1), Size: c.QueryInt("size", 10)}, &req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(res)
	}
}
