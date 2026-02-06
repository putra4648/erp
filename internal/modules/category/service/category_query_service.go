package service

import (
	categoryModel "putra4648/erp/internal/modules/category/model"
	categoryRepository "putra4648/erp/internal/modules/category/repository"

	"github.com/google/uuid"
)

type CategoryQueryService interface {
	GetCategoryByID(id uuid.UUID) (*categoryModel.CategoryResponse, error)
	GetAllCategories() ([]*categoryModel.CategoryResponse, error)
}

type categoryQueryService struct {
	categoryRepo categoryRepository.CategoryRepository
}

func NewCategoryQueryService(categoryRepo categoryRepository.CategoryRepository) CategoryQueryService {
	return &categoryQueryService{categoryRepo: categoryRepo}
}

func (s *categoryQueryService) GetCategoryByID(id uuid.UUID) (*categoryModel.CategoryResponse, error) {
	category, err := s.categoryRepo.FindByID(id)
	if err != nil {
		return nil, &CategoryError{Code: "NOT_FOUND", Message: "Category not found"}
	}

	return category.ToResponse(), nil
}

func (s *categoryQueryService) GetAllCategories() ([]*categoryModel.CategoryResponse, error) {
	categories, err := s.categoryRepo.FindAll()
	if err != nil {
		return nil, &CategoryError{Code: "DATABASE_ERROR", Message: "Failed to retrieve categories"}
	}

	responses := make([]*categoryModel.CategoryResponse, len(categories))
	for i, category := range categories {
		responses[i] = category.ToResponse()
	}

	return responses, nil
}
