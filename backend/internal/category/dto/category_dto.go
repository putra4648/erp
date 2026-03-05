package dto

type CategoryRequest struct {
	Name string `json:"name"`
	Page int    `json:"page"`
	Size int    `json:"size"`
}

type CategoryDTO struct {
	ID   string `json:"id"`
	Name string `json:"name" validate:"required,max=255"`
}
