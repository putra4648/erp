package service

import (
	"context"
	"putra4648/erp/internal/category/dto"
	"putra4648/erp/internal/category/mapper"
	"putra4648/erp/internal/category/repository"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type categoryCommandService struct {
	categoryRepo repository.CategoryRepository
	logger       *zap.Logger
}

func NewCategoryCommandService(categoryRepo repository.CategoryRepository, logger *zap.Logger) CategoryCommandService {
	return &categoryCommandService{
		categoryRepo: categoryRepo,
		logger:       logger,
	}
}

func (s *categoryCommandService) CreateCategory(ctx context.Context, categoryDTO *dto.CategoryDTO) (*dto.CategoryDTO, error) {
	category := mapper.ToCategory(categoryDTO)

	err := s.categoryRepo.Create(ctx, category)
	if err != nil {
		s.logger.Error("Failed to create category in DB", zap.Error(err), zap.String("name", category.Name))
		return nil, &CategoryError{Code: "DATABASE_ERROR", Message: "Failed to create category"}
	}

	return mapper.ToCategoryDTO(category), nil
}

func (s *categoryCommandService) UpdateCategory(ctx context.Context, id uuid.UUID, categoryDTO *dto.CategoryDTO) (*dto.CategoryDTO, error) {
	existingCategory, err := s.categoryRepo.FindByID(ctx, id)
	if err != nil {
		return nil, &CategoryError{Code: "NOT_FOUND", Message: "Category not found"}
	}

	existingCategory.Name = categoryDTO.Name

	err = s.categoryRepo.Update(ctx, existingCategory)
	if err != nil {
		s.logger.Error("Failed to update category in DB", zap.Error(err), zap.String("id", id.String()))
		return nil, &CategoryError{Code: "DATABASE_ERROR", Message: "Failed to update category"}
	}

	return mapper.ToCategoryDTO(existingCategory), nil
}

func (s *categoryCommandService) DeleteCategory(ctx context.Context, id uuid.UUID) error {
	_, err := s.categoryRepo.FindByID(ctx, id)
	if err != nil {
		return &CategoryError{Code: "NOT_FOUND", Message: "Category not found"}
	}

	err = s.categoryRepo.Delete(ctx, id)
	if err != nil {
		s.logger.Error("Failed to delete category in DB", zap.Error(err), zap.String("id", id.String()))
		return &CategoryError{Code: "DATABASE_ERROR", Message: "Failed to delete category"}
	}

	return nil
}
