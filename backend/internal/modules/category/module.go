package category

import (
	category_repository "putra4648/erp/internal/modules/category/repository"
	category_service "putra4648/erp/internal/modules/category/service"

	"go.uber.org/dig"
)

func Register(container *dig.Container) error {
	// Repository
	if err := container.Provide(category_repository.NewCategoryRepository); err != nil {
		return err
	}

	// Services
	if err := container.Provide(category_service.NewCategoryCommandService); err != nil {
		return err
	}
	if err := container.Provide(category_service.NewCategoryQueryService); err != nil {
		return err
	}

	return nil
}
