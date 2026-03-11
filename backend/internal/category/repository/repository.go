package repository

import (
	"context"
	"putra4648/erp/internal/category/domain"
	"putra4648/erp/internal/category/dto"
	sharedDto "putra4648/erp/internal/shared/dto"

	"github.com/google/uuid"
)

type CategoryRepository interface {
	Create(ctx context.Context, category *domain.Category) error
	Update(ctx context.Context, category *domain.Category) error
	Delete(ctx context.Context, id uuid.UUID) error

	FindByID(ctx context.Context, id uuid.UUID) (*domain.Category, error)
	FindAll(ctx context.Context, pagination *sharedDto.PaginationRequest, req *dto.CategoryDTO) ([]*domain.Category, int64, error)
}
