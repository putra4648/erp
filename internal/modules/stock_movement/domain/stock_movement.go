package domain

import (
	"putra4648/erp/internal/modules/shared/enums"
	. "putra4648/erp/internal/modules/shared/utils"
	domainWarehouse "putra4648/erp/internal/modules/warehouse/domain"
	"time"

	"github.com/google/uuid"
)

type StockMovement struct {
	ID                     uuid.UUID                  `gorm:"type:uuid;primarykey;default:uuid_generate_v4()"`
	MovementNo             string                     `gorm:"type:varchar(255);unique"`
	Type                   enums.TransferType         `gorm:"type:varchar(255)"` // IN, OUT, TRANSFER
	OriginWarehouseID      *uuid.UUID                 `gorm:"type:uuid;column:origin_warehouse_id"`
	OriginWarehouse        *domainWarehouse.Warehouse `gorm:"foreignKey:OriginWarehouseID"`
	DestinationWarehouseID *uuid.UUID                 `gorm:"type:uuid;column:destination_warehouse_id"`
	DestinationWarehouse   *domainWarehouse.Warehouse `gorm:"foreignKey:DestinationWarehouseID"`
	ReferenceNo            string                     `gorm:"type:varchar(255)"`
	Status                 enums.Status               `gorm:"type:varchar(255)"`
	TransactionDate        time.Time                  `gorm:"type:timestamp"`
	Note                   string                     `gorm:"type:text"`
	Items                  []StockMovementItem        `gorm:"foreignKey:StockMovementID"`
}

type StockMovementResponse struct {
	ID                     uuid.UUID                   `json:"id"`
	MovementNo             string                      `json:"movement_no"`
	Type                   string                      `json:"type"`
	OriginWarehouseID      *uuid.UUID                  `json:"origin_warehouse_id"`
	OriginWarehouse        *domainWarehouse.Warehouse  `json:"origin_warehouse,omitempty"`
	DestinationWarehouseID *uuid.UUID                  `json:"destination_warehouse_id"`
	DestinationWarehouse   *domainWarehouse.Warehouse  `json:"destination_warehouse,omitempty"`
	ReferenceNo            string                      `json:"reference_no"`
	Status                 enums.Status                `json:"status"`
	TransactionDate        time.Time                   `json:"transaction_date"`
	Note                   string                      `json:"note"`
	Items                  []StockMovementItemResponse `json:"items"`
}

type StockMovementDTO struct {
	ID                     string                 `json:"id"`
	MovementNo             string                 `json:"movement_no"`
	Type                   string                 `json:"type" validate:"required"`
	OriginWarehouseID      *string                `json:"origin_warehouse_id"`
	DestinationWarehouseID *string                `json:"destination_warehouse_id"`
	ReferenceNo            string                 `json:"reference_no"`
	Status                 enums.Status           `json:"status" default:"DRAFT"`
	TransactionDate        string                 `json:"transaction_date"`
	Note                   string                 `json:"note"`
	Items                  []StockMovementItemDTO `json:"items" validate:"required,min=1"`
}

func (s *StockMovement) ToResponse() *StockMovementResponse {
	return &StockMovementResponse{
		ID:                     s.ID,
		MovementNo:             s.MovementNo,
		Type:                   string(s.Type),
		OriginWarehouseID:      s.OriginWarehouseID,
		OriginWarehouse:        s.OriginWarehouse,
		DestinationWarehouseID: s.DestinationWarehouseID,
		DestinationWarehouse:   s.DestinationWarehouse,
		ReferenceNo:            s.ReferenceNo,
		Status:                 s.Status,
		TransactionDate:        s.TransactionDate,
		Note:                   s.Note,
		Items:                  MapSlice(s.Items, func(i StockMovementItem) StockMovementItemResponse { return *i.ToResponse() }),
	}
}

func (dto *StockMovementDTO) ToModel() *StockMovement {
	id, _ := uuid.Parse(dto.ID)
	var originID, destID *uuid.UUID

	if dto.OriginWarehouseID != nil && *dto.OriginWarehouseID != "" {
		parsed, _ := uuid.Parse(*dto.OriginWarehouseID)
		originID = &parsed
	}
	if dto.DestinationWarehouseID != nil && *dto.DestinationWarehouseID != "" {
		parsed, _ := uuid.Parse(*dto.DestinationWarehouseID)
		destID = &parsed
	}

	items := make([]StockMovementItem, len(dto.Items))
	for i, item := range dto.Items {
		pID, _ := uuid.Parse(item.ProductID)
		itemID, _ := uuid.Parse(item.ID)
		items[i] = StockMovementItem{
			ID:              itemID,
			ProductID:       pID,
			Quantity:        item.Quantity,
			Note:            item.Note,
			StockMovementID: id,
		}
	}

	transactionDate, _ := time.Parse("2006-01-02", dto.TransactionDate)

	return &StockMovement{
		ID:                     id,
		Type:                   enums.TransferType(dto.Type),
		OriginWarehouseID:      originID,
		DestinationWarehouseID: destID,
		ReferenceNo:            dto.ReferenceNo,
		Status:                 dto.Status,
		TransactionDate:        transactionDate,
		Note:                   dto.Note,
		Items:                  items,
	}
}
