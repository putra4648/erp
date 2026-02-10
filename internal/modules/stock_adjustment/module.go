package stock_adjustment

import (
	"putra4648/erp/internal/modules/stock_adjustment/repository"
	"putra4648/erp/internal/modules/stock_adjustment/service"

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
	if err := container.Provide(service.NewAdjustmentReasonService); err != nil {
		return err
	}
	if err := container.Provide(service.NewStockAdjustmentService); err != nil {
		return err
	}

	return nil
}
