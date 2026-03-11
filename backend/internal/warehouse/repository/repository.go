package repository

import (
	"context"
	sharedDto "putra4648/erp/internal/shared/dto"
	"putra4648/erp/internal/warehouse/domain"
	"putra4648/erp/internal/warehouse/dto"

	"github.com/google/uuid"
)

type WarehouseRepository interface {
	Save(ctx context.Context, warehouse *domain.Warehouse) error
	Update(ctx context.Context, warehouse *domain.Warehouse) error
	Delete(ctx context.Context, id uuid.UUID) error

	FindByID(ctx context.Context, id uuid.UUID) (*domain.Warehouse, error)
	FindAll(ctx context.Context, pagination *sharedDto.PaginationRequest, req *dto.WarehouseDTO) ([]*domain.Warehouse, int64, error)
}
