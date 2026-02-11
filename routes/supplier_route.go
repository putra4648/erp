package routes

import (
	supplierDto "putra4648/erp/internal/modules/supplier/dto"
	supplierService "putra4648/erp/internal/modules/supplier/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func RegisterSupplierRoutes(
	api fiber.Router,
	scs supplierService.SupplierCommandService,
	sqs supplierService.SupplierQueryService,
) {
	supplier := api.Group("/supplier")
	{
		supplier.Post("/", createSupplier(scs))
		supplier.Get("/:id", getSupplierByID(sqs))
		supplier.Get("/", getAllSuppliers(sqs))
		supplier.Put("/:id", updateSupplier(scs))
		supplier.Delete("/:id", deleteSupplier(scs))
	}
}

func createSupplier(service supplierService.SupplierCommandService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req supplierDto.CreateSupplierRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		id, err := service.Create(c.Context(), &req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id})
	}
}

func getSupplierByID(service supplierService.SupplierQueryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
		}
		supplier, err := service.FindByID(c.Context(), id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "supplier not found"})
		}
		return c.JSON(supplier)
	}
}

func getAllSuppliers(service supplierService.SupplierQueryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		page := c.QueryInt("page", 1)
		size := c.QueryInt("size", 10)
		name := c.Query("name")

		suppliers, err := service.FindAll(c.Context(), &supplierDto.SupplierFindAllRequest{Name: name, Page: page, Size: size})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(suppliers)
	}
}

func updateSupplier(service supplierService.SupplierCommandService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
		}
		var req supplierDto.UpdateSupplierRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		if err := service.Update(c.Context(), id, &req); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.SendStatus(fiber.StatusNoContent)
	}
}

func deleteSupplier(service supplierService.SupplierCommandService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
		}
		if err := service.Delete(c.Context(), id); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.SendStatus(fiber.StatusNoContent)
	}
}
