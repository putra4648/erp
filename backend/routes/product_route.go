package routes

import (
	"putra4648/erp/configs/logger"
	"putra4648/erp/internal/modules/product/dto"
	productService "putra4648/erp/internal/modules/product/service"
	. "putra4648/erp/internal/modules/shared/utils"

	"github.com/casbin/casbin/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func RegisterProductRoutes(
	app *fiber.App,
	api fiber.Router,
	productCommandService productService.ProductCommandService,
	productQueryService productService.ProductQueryService,
	enforcer *casbin.Enforcer,
) {
	// Product routes with Casbin authorization
	products := api.Group("/products")
	{
		products.Post("/", createProduct(productCommandService))
		products.Get("/:id", getProductByID(productQueryService))
		products.Get("/", getAllProducts(productQueryService))
		products.Put("/:id", updateProduct(productCommandService))
		products.Delete("/:id", deleteProduct(productCommandService))
	}
}

// Product handlers
func createProduct(service productService.ProductCommandService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req dto.ProductDTO
		if err := c.BodyParser(&req); err != nil {
			logger.Log.Error(err.Error())
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		// Create product
		response, err := service.CreateProduct(c.Context(), &req)
		if err != nil {
			if productErr, ok := err.(*productService.ProductError); ok {
				return c.Status(GetStatusCode(productErr.Code)).JSON(fiber.Map{"error": productErr.Message})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create product"})
		}

		return c.Status(fiber.StatusCreated).JSON(response)
	}
}

func getProductByID(service productService.ProductQueryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
		}

		// Get product by ID
		response, err := service.GetProductByID(c.Context(), id)
		if err != nil {
			if productErr, ok := err.(*productService.ProductError); ok {
				return c.Status(GetStatusCode(productErr.Code)).JSON(fiber.Map{"error": productErr.Message})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve product"})
		}

		return c.JSON(response)
	}
}

func getAllProducts(service productService.ProductQueryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		page := c.QueryInt("page", 1)
		size := c.QueryInt("size", 10)
		name := c.Query("name")

		// Get all products
		responses, err := service.GetAllProducts(c.Context(), &dto.ProductRequest{Name: name, Page: page, Size: size})
		if err != nil {
			if productErr, ok := err.(*productService.ProductError); ok {
				return c.Status(GetStatusCode(productErr.Code)).JSON(fiber.Map{"error": productErr.Message})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve products"})
		}

		return c.JSON(responses)
	}
}

func updateProduct(service productService.ProductCommandService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
		}

		var req dto.ProductDTO
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		// Update product
		response, err := service.UpdateProduct(c.Context(), id, &req)
		if err != nil {
			if productErr, ok := err.(*productService.ProductError); ok {
				return c.Status(GetStatusCode(productErr.Code)).JSON(fiber.Map{"error": productErr.Message})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update product"})
		}

		return c.JSON(response)
	}
}

func deleteProduct(service productService.ProductCommandService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
		}

		// Delete product
		err = service.DeleteProduct(c.Context(), id)
		if err != nil {
			if productErr, ok := err.(*productService.ProductError); ok {
				return c.Status(GetStatusCode(productErr.Code)).JSON(fiber.Map{"error": productErr.Message})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete product"})
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}
