package dto

import "github.com/shopspring/decimal"

type StockLevelRequest struct {
	ProductID   string `json:"product_id" query:"product_id"`
	WarehouseID string `json:"warehouse_id" query:"warehouse_id"`
	Search      string `json:"search" query:"search"`
	Page        int    `json:"page" query:"page"`
	Size        int    `json:"size" query:"size"`
}

type StockLevelResponse struct {
	ID            string          `json:"id"`
	ProductID     string          `json:"product_id"`
	ProductName   string          `json:"product_name"`
	WarehouseID   string          `json:"warehouse_id"`
	WarehouseName string          `json:"warehouse_name"`
	Quantity      decimal.Decimal `json:"quantity"`
	LastUpdated   string          `json:"last_updated"`
}
