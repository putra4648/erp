package dto

type UOMDTO struct {
	ID   string `json:"id"`
	Name string `json:"name" validate:"required,max=255"`
}
