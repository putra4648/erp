package routes

import (
	"putra4648/erp/configs/middleware"
	sharedDto "putra4648/erp/internal/shared/dto"
	"putra4648/erp/internal/stock_adjustment/dto"
	"putra4648/erp/internal/stock_adjustment/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func RegisterStockAdjustmentRoutes(
	api fiber.Router,
	saqs service.StockAdjustmentQueryService,
	sacs service.StockAdjustmentCommandService,
	arqs service.AdjustmentReasonQueryService,
	arcs service.AdjustmentReasonCommandService,
) {
	// Stock Adjustment routes
	adjustment := api.Group("/stock-adjustment")
	{
		adjustment.Post("/", middleware.RequirePermission("create:adjustments"), createStockAdjustment(sacs))
		adjustment.Get("/:id", middleware.RequirePermission("read:adjustments"), getStockAdjustmentByID(saqs))
		adjustment.Get("/", middleware.RequirePermission("read:adjustments"), getAllStockAdjustments(saqs))
		adjustment.Put("/:id", middleware.RequirePermission("update:adjustments"), updateStockAdjustment(sacs))
		adjustment.Post("/:id/approve", middleware.RequirePermission("approve:adjustments"), approveStockAdjustment(sacs))
		adjustment.Post("/:id/void", middleware.RequirePermission("void:adjustments"), voidStockAdjustment(sacs))
	}

	// Adjustment Reason routes
	reason := api.Group("/adjustment-reason")
	{
		reason.Post("/", middleware.RequirePermission("create:adjustment-reasons"), createAdjustmentReason(arcs))
		reason.Get("/", middleware.RequirePermission("read:adjustment-reasons"), getAllAdjustmentReasons(arqs))
	}
}

func createStockAdjustment(s service.StockAdjustmentCommandService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req dto.CreateStockAdjustmentRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		userIDStr := c.Locals("user_id").(string)
		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to parse user id"})
		}

		res, err := s.Create(c.Context(), &req, userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusCreated).JSON(res)
	}
}

func getStockAdjustmentByID(s service.StockAdjustmentQueryService) fiber.Handler {
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

func getAllStockAdjustments(s service.StockAdjustmentQueryService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		pagination := sharedDto.PaginationRequest{}
		pagination.Page = c.QueryInt("page", 1)
		pagination.Size = c.QueryInt("size", 10)

		req := dto.StockAdjustmentDto{}

		res, err := s.FindAll(c.Context(), &pagination, &req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(res)
	}
}

func createAdjustmentReason(s service.AdjustmentReasonCommandService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req dto.AdjustmentReasonRequest
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

func getAllAdjustmentReasons(s service.AdjustmentReasonQueryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		pagination := sharedDto.PaginationRequest{}
		pagination.Page = c.QueryInt("page", 1)
		pagination.Size = c.QueryInt("size", 10)

		req := dto.AdjustmentReasonDto{}

		res, err := s.FindAll(c.Context(), &pagination, &req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(res)
	}
}

func updateStockAdjustment(s service.StockAdjustmentCommandService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
		}

		var req dto.CreateStockAdjustmentRequest
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

func approveStockAdjustment(s service.StockAdjustmentCommandService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
		}

		userIDStr := c.Locals("user_id").(string)
		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to parse user id"})
		}

		res, err := s.Approve(c.Context(), id, userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(res)
	}
}

func voidStockAdjustment(s service.StockAdjustmentCommandService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
		}

		res, err := s.Void(c.Context(), id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(res)
	}
}
