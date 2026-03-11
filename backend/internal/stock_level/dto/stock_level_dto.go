package dto

import "github.com/shopspring/decimal"

type StockLevelDto struct {
	ProductID     *string         `json:"product_id" query:"product_id"`
	WarehouseID   *string         `json:"warehouse_id" query:"warehouse_id"`
	ID            string          `json:"id"`
	ProductName   string          `json:"product_name"`
	WarehouseName string          `json:"warehouse_name"`
	Quantity      decimal.Decimal `json:"quantity"`
	LastUpdated   string          `json:"last_updated"`
	Name          string          `json:"name" query:"name"`
}
