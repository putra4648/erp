package stock_movement

import (
	"putra4648/erp/internal/stock_movement/repository"
	"putra4648/erp/internal/stock_movement/service"

	"go.uber.org/dig"
)

func Register(container *dig.Container) error {
	if err := container.Provide(repository.NewStockMovementRepository); err != nil {
		return err
	}
	if err := container.Provide(service.NewStockMovementCommandService); err != nil {
		return err
	}
	if err := container.Provide(service.NewStockMovementQueryService); err != nil {
		return err
	}
	return nil
}
