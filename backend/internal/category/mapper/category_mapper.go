package mapper

import (
	"putra4648/erp/internal/category/domain"
	"putra4648/erp/internal/category/dto"

	"github.com/google/uuid"
)

func ToCategory(categoryDTO *dto.CategoryDTO) *domain.Category {
	return &domain.Category{
		Name: categoryDTO.Name,
	}
}

func ToCategories(categoryDTOs []*dto.CategoryDTO) []*domain.Category {
	categories := make([]*domain.Category, len(categoryDTOs))
	for i, categoryDTO := range categoryDTOs {
		var currentId uuid.UUID
		if categoryDTO.ID != "" {
			currentId = uuid.MustParse(categoryDTO.ID)
		} else {
			currentId = uuid.New()
		}
		categories[i] = &domain.Category{
			ID:   &currentId,
			Name: categoryDTO.Name,
		}
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
	categoryDTOs := make([]*dto.CategoryDTO, len(categories))
	for i, category := range categories {
		categoryDTOs[i] = &dto.CategoryDTO{
			ID:   category.ID.String(),
			Name: category.Name,
		}
	}
	return categoryDTOs
}
