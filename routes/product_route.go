package routes

import (
	productModel "putra4648/erp/internal/modules/product/model"
	productService "putra4648/erp/internal/modules/product/service"

	"github.com/casbin/casbin/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func RegisterProductRoutes(
	app *fiber.App,
	api fiber.Router,
	productCommandService productService.ProductCommandService,
	productQueryService productService.ProductQueryService,
	uomQueryService productService.UOMQueryService, // Added UOMQueryService
	enforcer *casbin.Enforcer,
) {
	// Product routes with Casbin authorization
	products := api.Group("/products")
	// products.Use(middleware.PermissionMiddleware(enforcer))
	{
		products.Post("/", createProduct(productCommandService))
		products.Get("/:id", getProductByID(productQueryService))
		products.Get("/", getAllProducts(productQueryService))
		products.Put("/:id", updateProduct(productCommandService))
		products.Delete("/:id", deleteProduct(productCommandService))
		products.Get("/uom", getAllUOMs(uomQueryService)) // Added new route for UOMs
	}
}

// Product handlers
func createProduct(service productService.ProductCommandService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req productModel.ProductDTO
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		// Create product
		response, err := service.CreateProduct(&req)
		if err != nil {
			if productErr, ok := err.(*productService.ProductError); ok {
				return c.Status(getStatusCode(productErr.Code)).JSON(fiber.Map{"error": productErr.Message})
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
		response, err := service.GetProductByID(id)
		if err != nil {
			if productErr, ok := err.(*productService.ProductError); ok {
				return c.Status(getStatusCode(productErr.Code)).JSON(fiber.Map{"error": productErr.Message})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve product"})
		}

		return c.JSON(response)
	}
}

func getAllProducts(service productService.ProductQueryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get all products
		responses, err := service.GetAllProducts()
		if err != nil {
			if productErr, ok := err.(*productService.ProductError); ok {
				return c.Status(getStatusCode(productErr.Code)).JSON(fiber.Map{"error": productErr.Message})
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

		var req productModel.ProductDTO
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		// Update product
		response, err := service.UpdateProduct(id, &req)
		if err != nil {
			if productErr, ok := err.(*productService.ProductError); ok {
				return c.Status(getStatusCode(productErr.Code)).JSON(fiber.Map{"error": productErr.Message})
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
		err = service.DeleteProduct(id)
		if err != nil {
			if productErr, ok := err.(*productService.ProductError); ok {
				return c.Status(getStatusCode(productErr.Code)).JSON(fiber.Map{"error": productErr.Message})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete product"})
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}

// UOM handlers
func getAllUOMs(service productService.UOMQueryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		responses, err := service.GetAllUOMs()
		if err != nil {
			if uomErr, ok := err.(*productService.UOMError); ok {
				return c.Status(getStatusCode(uomErr.Code)).JSON(fiber.Map{"error": uomErr.Message})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve UOMs"})
		}
		return c.JSON(responses)
	}
}

// Helper function to map error codes to HTTP status codes
func getStatusCode(code string) int {
	switch code {
	case "NOT_FOUND":
		return fiber.StatusNotFound
	case "DUPLICATE_SKU":
		return fiber.StatusConflict
	case "DATABASE_ERROR": // Added for UOM and general database errors
		return fiber.StatusInternalServerError
	default:
		return fiber.StatusInternalServerError
	}
}
