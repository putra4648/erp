package dto

type CategoryFindAllRequest struct {
	Name string `json:"name"`
	Page int    `json:"page"`
	Size int    `json:"size"`
}
