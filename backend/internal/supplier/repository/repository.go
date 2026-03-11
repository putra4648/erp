package repository

import (
	"context"
	sharedDto "putra4648/erp/internal/shared/dto"
	"putra4648/erp/internal/supplier/domain"
	"putra4648/erp/internal/supplier/dto"

	"github.com/google/uuid"
)

type SupplierRepository interface {
	Save(ctx context.Context, supplier *domain.Supplier) error
	Update(ctx context.Context, supplier *domain.Supplier) error
	Delete(ctx context.Context, id uuid.UUID) error

	FindByID(ctx context.Context, id uuid.UUID) (*domain.Supplier, error)
	FindAll(ctx context.Context, pagination *sharedDto.PaginationRequest, req *dto.SupplierDTO) ([]*domain.Supplier, int64, error)
}
