package stock_level

import (
	"putra4648/erp/internal/modules/stock_level/repository"
	"putra4648/erp/internal/modules/stock_level/service"

	"go.uber.org/dig"
)

func Register(container *dig.Container) error {
	if err := container.Provide(repository.NewStockLevelRepository); err != nil {
		return err
	}
	if err := container.Provide(service.NewStockLevelQueryService); err != nil {
		return err
	}
	if err := container.Provide(service.NewStockLevelCommandService); err != nil {
		return err
	}
	return nil
}
