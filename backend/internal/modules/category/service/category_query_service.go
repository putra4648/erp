package service

import (
	"context"
	"putra4648/erp/internal/modules/category/dto"
	"putra4648/erp/internal/modules/category/mapper"
	"putra4648/erp/internal/modules/category/repository"
	sharedDto "putra4648/erp/internal/modules/shared/dto"

	"github.com/google/uuid"
)

type categoryQueryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryQueryService(categoryRepo repository.CategoryRepository) CategoryQueryService {
	return &categoryQueryService{categoryRepo: categoryRepo}
}

func (s *categoryQueryService) GetCategoryByID(ctx context.Context, id uuid.UUID) (*dto.CategoryDTO, error) {
	category, err := s.categoryRepo.FindByID(ctx, id)
	if err != nil {
		return nil, &CategoryError{Code: "NOT_FOUND", Message: "Category not found"}
	}

	return mapper.ToCategoryDTO(category), nil
}

func (s *categoryQueryService) GetAllCategories(ctx context.Context, request *dto.CategoryRequest) (*sharedDto.PaginationResponse[*dto.CategoryDTO], error) {
	categories, total, err := s.categoryRepo.FindAll(ctx, request)
	if err != nil {
		return nil, &CategoryError{Code: "DATABASE_ERROR", Message: "Failed to retrieve categories"}
	}

	responses := mapper.ToCategoryDTOs(categories)

	return &sharedDto.PaginationResponse[*dto.CategoryDTO]{
		Items: responses,
		Total: total,
		Page:  request.Page,
		Size:  request.Size,
	}, nil
}
