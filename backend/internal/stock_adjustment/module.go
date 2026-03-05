package stock_adjustment

import (
	"putra4648/erp/internal/stock_adjustment/repository"
	"putra4648/erp/internal/stock_adjustment/service"

	"go.uber.org/dig"
)

func Register(container *dig.Container) error {
	// Repositories
	if err := container.Provide(repository.NewAdjustmentReasonRepository); err != nil {
		return err
	}
	if err := container.Provide(repository.NewStockAdjustmentRepository); err != nil {
		return err
	}

	// Services
	if err := container.Provide(service.NewAdjustmentReasonQueryService); err != nil {
		return err
	}
	if err := container.Provide(service.NewAdjustmentReasonCommandService); err != nil {
		return err
	}
	if err := container.Provide(service.NewStockAdjustmentQueryService); err != nil {
		return err
	}
	if err := container.Provide(service.NewStockAdjustmentCommandService); err != nil {
		return err
	}

	return nil
}
