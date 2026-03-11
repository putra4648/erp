package dto

type CategoryDTO struct {
	ID   string `json:"id"`
	Name string `json:"name" validate:"required,max=255"`
}
