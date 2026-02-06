package routes

import (
	uomModel "putra4648/erp/internal/modules/uom/model"
	uomService "putra4648/erp/internal/modules/uom/service"
	. "putra4648/erp/utils"

	"github.com/casbin/casbin/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func RegisterUOMRoutes(
	app *fiber.App,
	api fiber.Router,
	uomCommandService uomService.UOMCommandService,
	uomQueryService uomService.UOMQueryService,
	enforcer *casbin.Enforcer,
) {
	uoms := api.Group("/uoms")
	{
		uoms.Post("/", createUOM(uomCommandService))
		uoms.Get("/:id", getUOMByID(uomQueryService))
		uoms.Get("/", getAllUOMs(uomQueryService))
		uoms.Put("/:id", updateUOM(uomCommandService))
		uoms.Delete("/:id", deleteUOM(uomCommandService))
	}
}

func createUOM(service uomService.UOMCommandService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req uomModel.UOMDTO
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		response, err := service.CreateUOM(&req)
		if err != nil {
			if uomErr, ok := err.(*uomService.UOMError); ok {
				return c.Status(GetStatusCode(uomErr.Code)).JSON(fiber.Map{"error": uomErr.Message})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create UOM"})
		}

		return c.Status(fiber.StatusCreated).JSON(response)
	}
}

func getUOMByID(service uomService.UOMQueryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UOM ID"})
		}

		response, err := service.GetUOMByID(id)
		if err != nil {
			if uomErr, ok := err.(*uomService.UOMError); ok {
				return c.Status(GetStatusCode(uomErr.Code)).JSON(fiber.Map{"error": uomErr.Message})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve UOM"})
		}

		return c.JSON(response)
	}
}

func getAllUOMs(service uomService.UOMQueryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		responses, err := service.GetAllUOMs()
		if err != nil {
			if uomErr, ok := err.(*uomService.UOMError); ok {
				return c.Status(GetStatusCode(uomErr.Code)).JSON(fiber.Map{"error": uomErr.Message})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve UOMs"})
		}
		return c.JSON(responses)
	}
}

func updateUOM(service uomService.UOMCommandService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UOM ID"})
		}

		var req uomModel.UOMDTO
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		response, err := service.UpdateUOM(id, &req)
		if err != nil {
			if uomErr, ok := err.(*uomService.UOMError); ok {
				return c.Status(GetStatusCode(uomErr.Code)).JSON(fiber.Map{"error": uomErr.Message})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update UOM"})
		}

		return c.JSON(response)
	}
}

func deleteUOM(service uomService.UOMCommandService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UOM ID"})
		}

		err = service.DeleteUOM(id)
		if err != nil {
			if uomErr, ok := err.(*uomService.UOMError); ok {
				return c.Status(GetStatusCode(uomErr.Code)).JSON(fiber.Map{"error": uomErr.Message})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete UOM"})
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}
