package service

import (
	categoryModel "putra4648/erp/internal/modules/category/model"
	categoryRepository "putra4648/erp/internal/modules/category/repository"

	"github.com/google/uuid"
)

type CategoryCommandService interface {
	CreateCategory(categoryDTO *categoryModel.CategoryDTO) (*categoryModel.CategoryResponse, error)
	UpdateCategory(id uuid.UUID, categoryDTO *categoryModel.CategoryDTO) (*categoryModel.CategoryResponse, error)
	DeleteCategory(id uuid.UUID) error
}

type categoryCommandService struct {
	categoryRepo categoryRepository.CategoryRepository
}

func NewCategoryCommandService(categoryRepo categoryRepository.CategoryRepository) CategoryCommandService {
	return &categoryCommandService{categoryRepo: categoryRepo}
}

func (s *categoryCommandService) CreateCategory(categoryDTO *categoryModel.CategoryDTO) (*categoryModel.CategoryResponse, error) {
	category := categoryDTO.ToModel()

	err := s.categoryRepo.Create(category)
	if err != nil {
		return nil, &CategoryError{Code: "DATABASE_ERROR", Message: "Failed to create category"}
	}

	return category.ToResponse(), nil
}

func (s *categoryCommandService) UpdateCategory(id uuid.UUID, categoryDTO *categoryModel.CategoryDTO) (*categoryModel.CategoryResponse, error) {
	existingCategory, err := s.categoryRepo.FindByID(id)
	if err != nil {
		return nil, &CategoryError{Code: "NOT_FOUND", Message: "Category not found"}
	}

	existingCategory.Name = categoryDTO.Name

	err = s.categoryRepo.Update(existingCategory)
	if err != nil {
		return nil, &CategoryError{Code: "DATABASE_ERROR", Message: "Failed to update category"}
	}

	return existingCategory.ToResponse(), nil
}

func (s *categoryCommandService) DeleteCategory(id uuid.UUID) error {
	_, err := s.categoryRepo.FindByID(id)
	if err != nil {
		return &CategoryError{Code: "NOT_FOUND", Message: "Category not found"}
	}

	err = s.categoryRepo.Delete(id)
	if err != nil {
		return &CategoryError{Code: "DATABASE_ERROR", Message: "Failed to delete category"}
	}

	return nil
}
