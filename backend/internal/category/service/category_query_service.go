package service

import (
	"context"
	"putra4648/erp/internal/category/dto"
	"putra4648/erp/internal/category/mapper"
	"putra4648/erp/internal/category/repository"
	sharedDto "putra4648/erp/internal/shared/dto"
	"putra4648/erp/internal/shared/errors"

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
		return nil, &errors.ErrorDto{Code: "NOT_FOUND", Message: "Category not found"}
	}

	return mapper.ToCategoryDTO(category), nil
}

func (s *categoryQueryService) GetAllCategories(ctx context.Context, pagination *sharedDto.PaginationRequest, request *dto.CategoryDTO) (*sharedDto.PaginationResponse[*dto.CategoryDTO], error) {
	categories, total, err := s.categoryRepo.FindAll(ctx, pagination, request)
	if err != nil {
		return nil, &errors.ErrorDto{Code: "DATABASE_ERROR", Message: "Failed to retrieve categories"}
	}

	responses := mapper.ToCategoryDTOs(categories)

	return &sharedDto.PaginationResponse[*dto.CategoryDTO]{
		Items: responses,
		Total: total,
		Page:  pagination.Page,
		Size:  pagination.Size,
	}, nil
}
