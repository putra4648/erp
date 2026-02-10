package domain

import (
	domainProduct "putra4648/erp/internal/modules/inventory/stock/domain"
	domainWarehouse "putra4648/erp/internal/modules/inventory/warehouse/domain"
	"putra4648/erp/internal/modules/stock_adjustment/dto"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type StockAdjustment struct {
	ID              uuid.UUID                 `gorm:"type:uuid;primary_key" json:"id"`
	AdjustmentNo    string                    `gorm:"type:varchar(50);unique;not null" json:"adjustment_no"`
	WarehouseID     uuid.UUID                 `gorm:"type:uuid;not null" json:"warehouse_id"`
	Warehouse       domainWarehouse.Warehouse `gorm:"foreignKey:WarehouseID" json:"warehouse"`
	TransactionDate time.Time                 `gorm:"type:timestamp;not null" json:"transaction_date"`
	Status          string                    `gorm:"type:varchar(20);default:'DRAFT'" json:"status"`
	Note            string                    `gorm:"type:text" json:"note"`
	CreatedBy       uuid.UUID                 `gorm:"type:uuid;not null" json:"created_by"`
	ApprovedBy      *uuid.UUID                `gorm:"type:uuid" json:"approved_by"`
	Items           []StockAdjustmentItem     `gorm:"foreignKey:StockAdjustmentID" json:"items"`
}

type StockAdjustmentItem struct {
	ID                uuid.UUID             `gorm:"type:uuid;primary_key" json:"id"`
	StockAdjustmentID uuid.UUID             `gorm:"type:uuid;not null" json:"stock_adjustment_id"`
	ProductID         uuid.UUID             `gorm:"type:uuid;not null" json:"product_id"`
	Product           domainProduct.Product `gorm:"foreignKey:ProductID" json:"product"`
	ReasonID          uuid.UUID             `gorm:"type:uuid;not null" json:"reason_id"`
	Reason            AdjustmentReason      `gorm:"foreignKey:ReasonID" json:"reason"`
	ActualQty         decimal.Decimal       `gorm:"type:decimal(19,4)" json:"actual_qty"`
	SystemQty         decimal.Decimal       `gorm:"type:decimal(19,4)" json:"system_qty"`
	AdjustmentQty     decimal.Decimal       `gorm:"type:decimal(19,4)" json:"adjustment_qty"`
}

func (s *StockAdjustment) ToResponse() *dto.StockAdjustmentResponse {
	var items []dto.StockAdjustmentItemResponse
	for _, item := range s.Items {
		items = append(items, *item.ToResponse())
	}

	return &dto.StockAdjustmentResponse{
		ID:              s.ID,
		AdjustmentNo:    s.AdjustmentNo,
		WarehouseID:     s.WarehouseID,
		TransactionDate: s.TransactionDate,
		Status:          s.Status,
		Note:            s.Note,
		CreatedBy:       s.CreatedBy,
		ApprovedBy:      s.ApprovedBy,
		Items:           items,
	}
}

func (i *StockAdjustmentItem) ToResponse() *dto.StockAdjustmentItemResponse {
	return &dto.StockAdjustmentItemResponse{
		ID:            i.ID,
		ProductID:     i.ProductID,
		ProductName:   i.Product.Name,
		ReasonID:      i.ReasonID,
		ReasonName:    i.Reason.Name,
		ActualQty:     i.ActualQty,
		SystemQty:     i.SystemQty,
		AdjustmentQty: i.AdjustmentQty,
	}
}
