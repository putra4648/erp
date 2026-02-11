package service

import (
	"context"
	"putra4648/erp/internal/modules/category/domain"
	"putra4648/erp/internal/modules/category/dto"
	categoryRepository "putra4648/erp/internal/modules/category/repository"
	sharedDto "putra4648/erp/internal/modules/shared/dto"

	"github.com/google/uuid"
)

type CategoryQueryService interface {
	GetCategoryByID(ctx context.Context, id uuid.UUID) (*domain.CategoryResponse, error)
	GetAllCategories(ctx context.Context, request *dto.CategoryFindAllRequest) (*sharedDto.PaginationResponse[*domain.CategoryResponse], error)
}

type categoryQueryService struct {
	categoryRepo categoryRepository.CategoryRepository
}

func NewCategoryQueryService(categoryRepo categoryRepository.CategoryRepository) CategoryQueryService {
	return &categoryQueryService{categoryRepo: categoryRepo}
}

func (s *categoryQueryService) GetCategoryByID(ctx context.Context, id uuid.UUID) (*domain.CategoryResponse, error) {
	category, err := s.categoryRepo.FindByID(ctx, id)
	if err != nil {
		return nil, &CategoryError{Code: "NOT_FOUND", Message: "Category not found"}
	}

	return category.ToResponse(), nil
}

func (s *categoryQueryService) GetAllCategories(ctx context.Context, request *dto.CategoryFindAllRequest) (*sharedDto.PaginationResponse[*domain.CategoryResponse], error) {
	categories, total, err := s.categoryRepo.FindAll(ctx, request)
	if err != nil {
		return nil, &CategoryError{Code: "DATABASE_ERROR", Message: "Failed to retrieve categories"}
	}

	responses := make([]*domain.CategoryResponse, len(categories))
	for i, category := range categories {
		responses[i] = category.ToResponse()
	}

	return &sharedDto.PaginationResponse[*domain.CategoryResponse]{
		Items: responses,
		Total: total,
		Page:  request.Page,
		Size:  request.Size,
	}, nil
}
