package repository

import (
	"context"
	"putra4648/erp/internal/modules/supplier/domain"
	"putra4648/erp/internal/modules/supplier/dto"

	"github.com/google/uuid"
)

type SupplierRepository interface {
	Save(ctx context.Context, supplier *domain.Supplier) error
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Supplier, error)
	FindAll(ctx context.Context, req *dto.SupplierFindAllRequest) ([]*domain.Supplier, int64, error)
	Update(ctx context.Context, supplier *domain.Supplier) error
	Delete(ctx context.Context, id uuid.UUID) error
}
