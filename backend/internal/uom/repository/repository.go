package repository

import (
	"context"
	"putra4648/erp/internal/uom/domain"
	"putra4648/erp/internal/uom/dto"

	"github.com/google/uuid"
)

type UOMRepository interface {
	Create(ctx context.Context, uom *domain.UOM) error
	FindByID(ctx context.Context, id uuid.UUID) (*domain.UOM, error)
	FindAll(ctx context.Context, req *dto.UOMRequest) ([]*domain.UOM, int64, error)
	Update(ctx context.Context, uom *domain.UOM) error
	Delete(ctx context.Context, id uuid.UUID) error
}
