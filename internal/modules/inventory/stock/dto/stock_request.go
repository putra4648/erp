package dto

type StockRequest struct {
	ProductID string  `json:"product_id"`
	Qty       float64 `json:"qty"`
	Reason    string  `json:"reason"`
}
