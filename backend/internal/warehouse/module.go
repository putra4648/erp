package warehouse

import (
	"putra4648/erp/internal/warehouse/repository"
	"putra4648/erp/internal/warehouse/service"

	"go.uber.org/dig"
)

func Register(container *dig.Container) error {

	// Warehouse
	if err := container.Provide(repository.NewWarehouseRepository); err != nil {
		return err
	}
	if err := container.Provide(service.NewWarehouseCommandService); err != nil {
		return err
	}
	if err := container.Provide(service.NewWarehouseQueryService); err != nil {
		return err
	}

	return nil
}
