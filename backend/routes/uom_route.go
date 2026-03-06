package routes

import (
	"putra4648/erp/configs/middleware"
	. "putra4648/erp/internal/shared/utils"
	"putra4648/erp/internal/uom/dto"
	"putra4648/erp/internal/uom/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func RegisterUOMRoutes(
	app *fiber.App,
	api fiber.Router,
	uomCommandService service.UOMCommandService,
	uomQueryService service.UOMQueryService,
) {
	uoms := api.Group("/uoms")
	{
		uoms.Post("/", middleware.RequirePermission("create:uoms"), createUOM(uomCommandService))
		uoms.Get("/:id", middleware.RequirePermission("read:uoms"), getUOMByID(uomQueryService))
		uoms.Get("/", middleware.RequirePermission("read:uoms"), getAllUOMs(uomQueryService))
		uoms.Put("/:id", middleware.RequirePermission("update:uoms"), updateUOM(uomCommandService))
		uoms.Delete("/:id", middleware.RequirePermission("delete:uoms"), deleteUOM(uomCommandService))
	}
}

func createUOM(s service.UOMCommandService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req dto.UOMDTO
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		response, err := s.CreateUOM(c.Context(), &req)
		if err != nil {
			if uomErr, ok := err.(*service.UOMError); ok {
				return c.Status(GetStatusCode(uomErr.Code)).JSON(fiber.Map{"error": uomErr.Message})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create UOM"})
		}

		return c.Status(fiber.StatusCreated).JSON(response)
	}
}

func getUOMByID(s service.UOMQueryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UOM ID"})
		}

		response, err := s.GetUOMByID(c.Context(), id)
		if err != nil {
			if uomErr, ok := err.(*service.UOMError); ok {
				return c.Status(GetStatusCode(uomErr.Code)).JSON(fiber.Map{"error": uomErr.Message})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve UOM"})
		}

		return c.JSON(response)
	}
}

func getAllUOMs(s service.UOMQueryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.Query("name")
		page := c.QueryInt("page", 1)
		size := c.QueryInt("size", 10)

		responses, err := s.GetAllUOMs(c.Context(), &dto.UOMRequest{Name: name, Page: page, Size: size})
		if err != nil {
			if uomErr, ok := err.(*service.UOMError); ok {
				return c.Status(GetStatusCode(uomErr.Code)).JSON(fiber.Map{"error": uomErr.Message})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve UOMs"})
		}
		return c.JSON(responses)
	}
}

func updateUOM(s service.UOMCommandService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UOM ID"})
		}

		var req dto.UOMDTO
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		response, err := s.UpdateUOM(c.Context(), id, &req)
		if err != nil {
			if uomErr, ok := err.(*service.UOMError); ok {
				return c.Status(GetStatusCode(uomErr.Code)).JSON(fiber.Map{"error": uomErr.Message})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update UOM"})
		}

		return c.JSON(response)
	}
}

func deleteUOM(s service.UOMCommandService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UOM ID"})
		}

		err = s.DeleteUOM(c.Context(), id)
		if err != nil {
			if uomErr, ok := err.(*service.UOMError); ok {
				return c.Status(GetStatusCode(uomErr.Code)).JSON(fiber.Map{"error": uomErr.Message})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete UOM"})
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}
