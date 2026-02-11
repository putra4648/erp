package routes

import (
	"putra4648/erp/internal/modules/stock_adjustment/dto"
	"putra4648/erp/internal/modules/stock_adjustment/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func RegisterStockAdjustmentRoutes(
	api fiber.Router,
	sas service.StockAdjustmentService,
	ars service.AdjustmentReasonService,
) {
	// Stock Adjustment routes
	adjustment := api.Group("/stock-adjustment")
	{
		adjustment.Post("/", createStockAdjustment(sas))
		adjustment.Get("/:id", getStockAdjustmentByID(sas))
		adjustment.Get("/", getAllStockAdjustments(sas))
		adjustment.Put("/:id", updateStockAdjustment(sas))
		adjustment.Post("/:id/approve", approveStockAdjustment(sas))
		adjustment.Post("/:id/void", voidStockAdjustment(sas))
	}

	// Adjustment Reason routes
	reason := api.Group("/adjustment-reason")
	{
		reason.Post("/", createAdjustmentReason(ars))
		reason.Get("/", getAllAdjustmentReasons(ars))
	}
}

func createStockAdjustment(s service.StockAdjustmentService) fiber.Handler {
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

func getStockAdjustmentByID(s service.StockAdjustmentService) fiber.Handler {
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

func getAllStockAdjustments(s service.StockAdjustmentService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		page := c.QueryInt("page", 1)
		size := c.QueryInt("size", 10)
		res, total, err := s.FindAll(c.Context(), page, size)
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

func createAdjustmentReason(s service.AdjustmentReasonService) fiber.Handler {
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

func getAllAdjustmentReasons(s service.AdjustmentReasonService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		res, err := s.FindAll(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(res)
	}
}

func updateStockAdjustment(s service.StockAdjustmentService) fiber.Handler {
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

func approveStockAdjustment(s service.StockAdjustmentService) fiber.Handler {
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

func voidStockAdjustment(s service.StockAdjustmentService) fiber.Handler {
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
