package dto

type ProductRequest struct {
	Name string `json:"name"`
	Page int    `json:"page"`
	Size int    `json:"size"`
}
