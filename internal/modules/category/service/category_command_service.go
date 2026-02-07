package service

import (
	"context"
	"putra4648/erp/internal/modules/category/domain"
	"putra4648/erp/internal/modules/category/repository"

	"github.com/google/uuid"
)

type CategoryCommandService interface {
	CreateCategory(ctx context.Context, categoryDTO *domain.CategoryDTO) (*domain.CategoryResponse, error)
	UpdateCategory(ctx context.Context, id uuid.UUID, categoryDTO *domain.CategoryDTO) (*domain.CategoryResponse, error)
	DeleteCategory(ctx context.Context, id uuid.UUID) error
}

type categoryCommandService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryCommandService(categoryRepo repository.CategoryRepository) CategoryCommandService {
	return &categoryCommandService{categoryRepo: categoryRepo}
}

func (s *categoryCommandService) CreateCategory(ctx context.Context, categoryDTO *domain.CategoryDTO) (*domain.CategoryResponse, error) {
	category := categoryDTO.ToModel()

	err := s.categoryRepo.Create(ctx, category)
	if err != nil {
		return nil, &CategoryError{Code: "DATABASE_ERROR", Message: "Failed to create category"}
	}

	return category.ToResponse(), nil
}

func (s *categoryCommandService) UpdateCategory(ctx context.Context, id uuid.UUID, categoryDTO *domain.CategoryDTO) (*domain.CategoryResponse, error) {
	existingCategory, err := s.categoryRepo.FindByID(ctx, id)
	if err != nil {
		return nil, &CategoryError{Code: "NOT_FOUND", Message: "Category not found"}
	}

	existingCategory.Name = categoryDTO.Name

	err = s.categoryRepo.Update(ctx, existingCategory)
	if err != nil {
		return nil, &CategoryError{Code: "DATABASE_ERROR", Message: "Failed to update category"}
	}

	return existingCategory.ToResponse(), nil
}

func (s *categoryCommandService) DeleteCategory(ctx context.Context, id uuid.UUID) error {
	_, err := s.categoryRepo.FindByID(ctx, id)
	if err != nil {
		return &CategoryError{Code: "NOT_FOUND", Message: "Category not found"}
	}

	err = s.categoryRepo.Delete(ctx, id)
	if err != nil {
		return &CategoryError{Code: "DATABASE_ERROR", Message: "Failed to delete category"}
	}

	return nil
}
