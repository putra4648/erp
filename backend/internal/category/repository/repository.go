package repository

import (
	"context"
	"putra4648/erp/internal/category/dto"

	"github.com/google/uuid"

	"putra4648/erp/internal/category/domain"
)

type CategoryRepository interface {
	Create(ctx context.Context, category *domain.Category) error
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Category, error)
	FindAll(ctx context.Context, req *dto.CategoryRequest) ([]*domain.Category, int64, error)
	Update(ctx context.Context, category *domain.Category) error
	Delete(ctx context.Context, id uuid.UUID) error
}
