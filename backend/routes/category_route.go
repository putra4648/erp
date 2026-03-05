package routes

import (
	"putra4648/erp/internal/category/dto"
	categoryDto "putra4648/erp/internal/category/dto"
	categoryService "putra4648/erp/internal/category/service"
	. "putra4648/erp/internal/shared/utils"

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
		var req categoryDto.CategoryDTO
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		response, err := service.CreateCategory(c.Context(), &req)
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

		response, err := service.GetCategoryByID(c.Context(), id)
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
		page := c.QueryInt("page", 1)
		size := c.QueryInt("size", 10)
		name := c.Query("name")

		responses, err := service.GetAllCategories(c.Context(), &dto.CategoryRequest{Name: name, Page: page, Size: size})
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

		var req categoryDto.CategoryDTO
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		response, err := service.UpdateCategory(c.Context(), id, &req)
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

		err = service.DeleteCategory(c.Context(), id)
		if err != nil {
			if categoryErr, ok := err.(*categoryService.CategoryError); ok {
				return c.Status(GetStatusCode(categoryErr.Code)).JSON(fiber.Map{"error": categoryErr.Message})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete category"})
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}
