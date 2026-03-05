package dto

type PaginationRequest struct {
	Page int `query:"page"`
	Size int `query:"size"`
}

type PaginationResponse[T any] struct {
	Items      []T   `json:"items"`
	Total      int64 `json:"total"`
	Page       int   `json:"page"`
	Size       int   `json:"size"`
}
