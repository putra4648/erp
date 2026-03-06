package routes

import (
	"putra4648/erp/configs/middleware"
	warehouseDto "putra4648/erp/internal/warehouse/dto"
	warehouseService "putra4648/erp/internal/warehouse/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func RegisterWarehouseRoutes(
	api fiber.Router,
	wcs warehouseService.WarehouseCommandService,
	wqs warehouseService.WarehouseQueryService,
) {
	warehouse := api.Group("/warehouse")
	{
		warehouse.Post("/", middleware.RequirePermission("create:warehouses"), createWarehouse(wcs))
		warehouse.Get("/:id", middleware.RequirePermission("read:warehouses"), getWarehouseByID(wqs))
		warehouse.Get("/", middleware.RequirePermission("read:warehouses"), getAllWarehouses(wqs))
		warehouse.Put("/:id", middleware.RequirePermission("update:warehouses"), updateWarehouse(wcs))
		warehouse.Delete("/:id", middleware.RequirePermission("delete:warehouses"), deleteWarehouse(wcs))
	}
}

func createWarehouse(service warehouseService.WarehouseCommandService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req warehouseDto.WarehouseDto
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

func getWarehouseByID(service warehouseService.WarehouseQueryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
		}
		warehouse, err := service.FindByID(c.Context(), id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "warehouse not found"})
		}
		return c.JSON(warehouse)
	}
}

func getAllWarehouses(service warehouseService.WarehouseQueryService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		page := c.QueryInt("page", 1)
		size := c.QueryInt("size", 10)
		name := c.Query("name")

		warehouses, err := service.FindAll(c.Context(), &warehouseDto.WarehouseFindAllRequest{Name: name, Page: page, Size: size})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(warehouses)
	}
}

func updateWarehouse(service warehouseService.WarehouseCommandService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
		}
		var req warehouseDto.WarehouseDto
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

func deleteWarehouse(service warehouseService.WarehouseCommandService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
		}
		dto, err := service.Delete(c.Context(), id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(dto)
	}
}
