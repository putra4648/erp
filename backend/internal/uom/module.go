package uom

import (
	uom_repository "putra4648/erp/internal/uom/repository"
	uom_service "putra4648/erp/internal/uom/service"

	"go.uber.org/dig"
)

func Register(container *dig.Container) error {
	// Repository
	if err := container.Provide(uom_repository.NewUOMRepository); err != nil {
		return err
	}

	// Services
	if err := container.Provide(uom_service.NewUOMCommandService); err != nil {
		return err
	}
	if err := container.Provide(uom_service.NewUOMQueryService); err != nil {
		return err
	}

	return nil
}
