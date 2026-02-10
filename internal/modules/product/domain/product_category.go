package domain

import (
	categoryDomain "putra4648/erp/internal/modules/category/domain"

	"github.com/google/uuid"
)

type ProductCategory struct {
	ProductID  uuid.UUID               `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	CategoryID uuid.UUID               `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Category   categoryDomain.Category `gorm:"foreignKey:CategoryID"`
}

type ProductCategoryResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (pc *ProductCategory) ToResponse() *ProductCategoryResponse {
	return &ProductCategoryResponse{
		ID:   pc.CategoryID,
		Name: pc.Category.Name,
	}
}

type ProductCategoryDTO struct {
	ID uuid.UUID `json:"id"`
}

func (pc *ProductCategory) ToDTO() *ProductCategoryDTO {
	return &ProductCategoryDTO{
		ID: pc.CategoryID,
	}
}

func (pc *ProductCategoryDTO) ToModel() *ProductCategory {
	return &ProductCategory{
		CategoryID: pc.ID,
	}
}
