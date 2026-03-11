package repository

import (
	"context"
	sharedDto "putra4648/erp/internal/shared/dto"
	"putra4648/erp/internal/uom/domain"
	"putra4648/erp/internal/uom/dto"

	"github.com/google/uuid"
)

type UOMRepository interface {
	Create(ctx context.Context, uom *domain.UOM) error
	Update(ctx context.Context, uom *domain.UOM) error
	Delete(ctx context.Context, id uuid.UUID) error

	FindByID(ctx context.Context, id uuid.UUID) (*domain.UOM, error)
	FindAll(ctx context.Context, pagination *sharedDto.PaginationRequest, req *dto.UOMDTO) ([]*domain.UOM, int64, error)
}
