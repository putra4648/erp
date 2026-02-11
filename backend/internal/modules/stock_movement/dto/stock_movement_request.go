package dto

type StockMovementRequest struct {
	Type string `json:"type"`
	Page int    `json:"page"`
	Size int    `json:"size"`
}
