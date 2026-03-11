package dto

import (
	"putra4648/erp/internal/shared/enums"

	"github.com/shopspring/decimal"
)

type StockMovementItemDTO struct {
	ID        string          `json:"id"`
	ProductID string          `json:"product_id" validate:"required"`
	Quantity  decimal.Decimal `json:"quantity" validate:"required,gt=0"`
	Note      string          `json:"note"`
}

type StockMovementDTO struct {
	ID                     string                  `json:"id"`
	MovementNo             string                  `json:"movement_no"`
	OriginWarehouseID      string                  `json:"origin_warehouse_id"`
	DestinationWarehouseID string                  `json:"destination_warehouse_id"`
	ReferenceNo            string                  `json:"reference_no"`
	Status                 enums.Status            `json:"status" default:"DRAFT"`
	TransactionDate        string                  `json:"transaction_date"`
	Note                   string                  `json:"note"`
	Items                  []*StockMovementItemDTO `json:"items" validate:"required,min=1"`
	Type                   string                  `json:"type" query:"type"`
	Name                   string                  `json:"name" query:"name"`
}
