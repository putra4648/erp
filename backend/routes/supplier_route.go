package routes

import (
	"putra4648/erp/configs/middleware"
	sharedDto "putra4648/erp/internal/shared/dto"
	"putra4648/erp/internal/supplier/dto"
	supplierService "putra4648/erp/internal/supplier/service"

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
		supplier.Post("/", middleware.RequirePermission("create:suppliers"), createSupplier(scs))
		supplier.Get("/:id", middleware.RequirePermission("read:suppliers"), getSupplierByID(sqs))
		supplier.Get("/", middleware.RequirePermission("read:suppliers"), getAllSuppliers(sqs))
		supplier.Put("/:id", middleware.RequirePermission("update:suppliers"), updateSupplier(scs))
		supplier.Delete("/:id", middleware.RequirePermission("delete:suppliers"), deleteSupplier(scs))
	}
}

func createSupplier(service supplierService.SupplierCommandService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req dto.SupplierDTO
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
		supplier, err := service.FindByID(c.Context(), id.String())
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "supplier not found"})
		}
		return c.JSON(supplier)
	}
}

func getAllSuppliers(service supplierService.SupplierQueryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req dto.SupplierDTO
		if err := c.QueryParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		suppliers, err := service.FindAll(c.Context(), &sharedDto.PaginationRequest{Page: c.QueryInt("page", 1), Size: c.QueryInt("size", 10)}, &req)
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
		var req dto.SupplierDTO
		req.ID = id.String()
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		dto, err := service.Update(c.Context(), &req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(dto)
	}
}

func deleteSupplier(service supplierService.SupplierCommandService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
		}
		dto, err := service.Delete(c.Context(), id.String())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(dto)
	}
}
