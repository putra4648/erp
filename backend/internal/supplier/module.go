package supplier

import (
	"putra4648/erp/internal/supplier/repository"
	"putra4648/erp/internal/supplier/service"

	"go.uber.org/dig"
)

func Register(container *dig.Container) error {

	// Supplier
	if err := container.Provide(repository.NewSupplierRepository); err != nil {
		return err
	}
	if err := container.Provide(service.NewSupplierCommandService); err != nil {
		return err
	}
	if err := container.Provide(service.NewSupplierQueryService); err != nil {
		return err
	}

	return nil
}
