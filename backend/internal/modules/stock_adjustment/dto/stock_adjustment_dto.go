package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CreateStockAdjustmentRequest struct {
	WarehouseID     uuid.UUID                          `json:"warehouse_id" validate:"required"`
	TransactionDate string                             `json:"transaction_date" validate:"required"`
	Note            string                             `json:"note"`
	Items           []CreateStockAdjustmentItemRequest `json:"items" validate:"required,dive,required"`
}

type CreateStockAdjustmentItemRequest struct {
	ProductID uuid.UUID       `json:"product_id" validate:"required"`
	ReasonID  uuid.UUID       `json:"reason_id" validate:"required"`
	ActualQty decimal.Decimal `json:"actual_qty"`
	SystemQty decimal.Decimal `json:"system_qty"`
}

type UpdateStockAdjustmentRequest struct {
	Note  string                             `json:"note"`
	Items []CreateStockAdjustmentItemRequest `json:"items" validate:"dive,required"`
}

type StockAdjustmentDto struct {
	ID              uuid.UUID                `json:"id"`
	AdjustmentNo    string                   `json:"adjustment_no"`
	WarehouseID     uuid.UUID                `json:"warehouse_id"`
	TransactionDate time.Time                `json:"transaction_date"`
	Status          string                   `json:"status"`
	Note            string                   `json:"note"`
	CreatedBy       uuid.UUID                `json:"created_by"`
	ApprovedBy      *uuid.UUID               `json:"approved_by,omitempty"`
	Items           []StockAdjustmentItemDto `json:"items"`
}

type StockAdjustmentItemDto struct {
	ID            uuid.UUID       `json:"id"`
	ProductID     uuid.UUID       `json:"product_id"`
	ProductName   string          `json:"product_name"`
	ReasonID      uuid.UUID       `json:"reason_id"`
	ReasonName    string          `json:"reason_name"`
	ActualQty     decimal.Decimal `json:"actual_qty"`
	SystemQty     decimal.Decimal `json:"system_qty"`
	AdjustmentQty decimal.Decimal `json:"adjustment_qty"`
}
