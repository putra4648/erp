package mapper

import (
	"putra4648/erp/internal/modules/category/domain"
	"putra4648/erp/internal/modules/category/dto"
)

func ToCategory(categoryDTO *dto.CategoryDTO) *domain.Category {
	return &domain.Category{
		Name: categoryDTO.Name,
	}
}

func ToCategories(categoryDTOs []*dto.CategoryDTO) []*domain.Category {
	var categories []*domain.Category
	for _, categoryDTO := range categoryDTOs {
		categories = append(categories, &domain.Category{
			Name: categoryDTO.Name,
		})
	}
	return categories
}

func ToCategoryDTO(category *domain.Category) *dto.CategoryDTO {
	return &dto.CategoryDTO{
		ID:   category.ID.String(),
		Name: category.Name,
	}
}

func ToCategoryDTOs(categories []*domain.Category) []*dto.CategoryDTO {
	var categoryDTOs []*dto.CategoryDTO
	for _, category := range categories {
		categoryDTOs = append(categoryDTOs, &dto.CategoryDTO{
			ID:   category.ID.String(),
			Name: category.Name,
		})
	}
	return categoryDTOs
}
