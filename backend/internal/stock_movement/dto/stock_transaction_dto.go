package dto

import (
	"github.com/shopspring/decimal"
)

type StockTransactionRequest struct {
	ProductID   string `query:"product_id"`
	WarehouseID string `query:"warehouse_id"`
	Page        int    `query:"page"`
	Size        int    `query:"size"`
}

type StockTransactionResponse struct {
	ID            string          `json:"id"`
	ProductID     string          `json:"product_id"`
	ProductName   string          `json:"product_name"`
	WarehouseID   string          `json:"warehouse_id"`
	WarehouseName string          `json:"warehouse_name"`
	SupplierID    string          `json:"supplier_id"`
	SupplierName  string          `json:"supplier_name"`
	Type          string          `json:"type"`
	Quantity      decimal.Decimal `json:"quantity"`
	ReferenceNo   string          `json:"reference_no"`
	CreatedAt     string          `json:"created_at"`
}
