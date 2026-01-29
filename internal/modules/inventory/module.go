package inventory

import (
	supplier_repository "putra4648/erp/internal/modules/inventory/supplier/repository"
	supplier_service "putra4648/erp/internal/modules/inventory/supplier/service"
	warehouse_repository "putra4648/erp/internal/modules/inventory/warehouse/repository"
	warehouse_service "putra4648/erp/internal/modules/inventory/warehouse/service"

	"go.uber.org/dig"
)

func Register(container *dig.Container) error {
	// Warehouse
	if err := container.Provide(warehouse_repository.NewWarehouseRepository); err != nil {
		return err
	}
	if err := container.Provide(warehouse_service.NewWarehouseCommandService); err != nil {
		return err
	}
	if err := container.Provide(warehouse_service.NewWarehouseQueryService); err != nil {
		return err
	}

	// Supplier
	if err := container.Provide(supplier_repository.NewSupplierRepository); err != nil {
		return err
	}
	if err := container.Provide(supplier_service.NewSupplierCommandService); err != nil {
		return err
	}
	if err := container.Provide(supplier_service.NewSupplierQueryService); err != nil {
		return err
	}

	return nil
}