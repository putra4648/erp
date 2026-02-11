package domain

import "github.com/google/uuid"

type Category struct {
	ID   uuid.UUID `gorm:"type:uuid;primarykey;default:uuid_generate_v4()"`
	Name string    `gorm:"type:varchar(255)"`
}

type CategoryResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type CategoryDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (dto *CategoryDTO) ToModel() *Category {
	id, _ := uuid.Parse(dto.ID)
	return &Category{
		ID:   id,
		Name: dto.Name,
	}
}

func (c *Category) ToResponse() *CategoryResponse {
	return &CategoryResponse{
		ID:   c.ID,
		Name: c.Name,
	}
}
