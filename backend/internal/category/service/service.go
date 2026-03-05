package service

import (
	"context"
	"putra4648/erp/internal/category/dto"
	sharedDto "putra4648/erp/internal/shared/dto"

	"github.com/google/uuid"
)

type CategoryQueryService interface {
	GetCategoryByID(ctx context.Context, id uuid.UUID) (*dto.CategoryDTO, error)
	GetAllCategories(ctx context.Context, request *dto.CategoryRequest) (*sharedDto.PaginationResponse[*dto.CategoryDTO], error)
}

type CategoryCommandService interface {
	CreateCategory(ctx context.Context, categoryDTO *dto.CategoryDTO) (*dto.CategoryDTO, error)
	UpdateCategory(ctx context.Context, id uuid.UUID, categoryDTO *dto.CategoryDTO) (*dto.CategoryDTO, error)
	DeleteCategory(ctx context.Context, id uuid.UUID) error
}
