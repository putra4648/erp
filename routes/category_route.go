package routes

import (
	categoryModel "putra4648/erp/internal/modules/category/model"
	categoryService "putra4648/erp/internal/modules/category/service"
	. "putra4648/erp/utils"

	"github.com/casbin/casbin/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func RegisterCategoryRoutes(
	app *fiber.App,
	api fiber.Router,
	categoryCommandService categoryService.CategoryCommandService,
	categoryQueryService categoryService.CategoryQueryService,
	enforcer *casbin.Enforcer,
) {
	categories := api.Group("/categories")
	{
		categories.Post("/", createCategory(categoryCommandService))
		categories.Get("/:id", getCategoryByID(categoryQueryService))
		categories.Get("/", getAllCategories(categoryQueryService))
		categories.Put("/:id", updateCategory(categoryCommandService))
		categories.Delete("/:id", deleteCategory(categoryCommandService))
	}
}

func createCategory(service categoryService.CategoryCommandService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req categoryModel.CategoryDTO
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		response, err := service.CreateCategory(&req)
		if err != nil {
			if categoryErr, ok := err.(*categoryService.CategoryError); ok {
				return c.Status(GetStatusCode(categoryErr.Code)).JSON(fiber.Map{"error": categoryErr.Message})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create category"})
		}

		return c.Status(fiber.StatusCreated).JSON(response)
	}
}

func getCategoryByID(service categoryService.CategoryQueryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid category ID"})
		}

		response, err := service.GetCategoryByID(id)
		if err != nil {
			if categoryErr, ok := err.(*categoryService.CategoryError); ok {
				return c.Status(GetStatusCode(categoryErr.Code)).JSON(fiber.Map{"error": categoryErr.Message})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve category"})
		}

		return c.JSON(response)
	}
}

func getAllCategories(service categoryService.CategoryQueryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		responses, err := service.GetAllCategories()
		if err != nil {
			if categoryErr, ok := err.(*categoryService.CategoryError); ok {
				return c.Status(GetStatusCode(categoryErr.Code)).JSON(fiber.Map{"error": categoryErr.Message})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve categories"})
		}

		return c.JSON(responses)
	}
}

func updateCategory(service categoryService.CategoryCommandService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid category ID"})
		}

		var req categoryModel.CategoryDTO
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		response, err := service.UpdateCategory(id, &req)
		if err != nil {
			if categoryErr, ok := err.(*categoryService.CategoryError); ok {
				return c.Status(GetStatusCode(categoryErr.Code)).JSON(fiber.Map{"error": categoryErr.Message})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update category"})
		}

		return c.JSON(response)
	}
}

func deleteCategory(service categoryService.CategoryCommandService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid category ID"})
		}

		err = service.DeleteCategory(id)
		if err != nil {
			if categoryErr, ok := err.(*categoryService.CategoryError); ok {
				return c.Status(GetStatusCode(categoryErr.Code)).JSON(fiber.Map{"error": categoryErr.Message})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete category"})
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}
