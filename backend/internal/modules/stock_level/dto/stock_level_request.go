package dto

type StockLevelRequest struct {
	ProductID   string `json:"product_id"`
	WarehouseID string `json:"warehouse_id"`
	Quantity    int    `json:"quantity"`
}
