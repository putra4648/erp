package model

import (
	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	ProductID uuid.UUID
	Name      string `gorm:"not null;size:255"`
}

type CategoryDTO struct {
	Name string `json:"name" validate:"required,max=255"`
}

type CategoryResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (c *Category) ToResponse() *CategoryResponse {
	return &CategoryResponse{
		ID:   c.ID,
		Name: c.Name,
	}
}

func (dto *CategoryDTO) ToModel() *Category {
	return &Category{
		Name: dto.Name,
	}
}
