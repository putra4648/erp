package product

import (
	product_repository "putra4648/erp/internal/product/repository"
	product_service "putra4648/erp/internal/product/service"

	"go.uber.org/dig"
)

func Register(container *dig.Container) error {
	// Repositories
	if err := container.Provide(product_repository.NewProductRepository); err != nil {
		return err
	}

	// Services
	if err := container.Provide(product_service.NewProductCommandService); err != nil {
		return err
	}
	if err := container.Provide(product_service.NewProductQueryService); err != nil {
		return err
	}

	return nil
}
