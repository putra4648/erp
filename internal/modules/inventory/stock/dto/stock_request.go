package dto

type StockRequest struct {
	ProductID string `json:"product_id"`
	Qty       string `json:"qty"`
	Reason    string `json:"reason"`
}
