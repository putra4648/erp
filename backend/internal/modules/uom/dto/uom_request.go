package dto

type UOMRequest struct {
	Name string `json:"name"`
	Page int    `json:"page"`
	Size int    `json:"size"`
}
