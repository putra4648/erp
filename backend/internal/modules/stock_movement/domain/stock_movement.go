package domain

import (
	"putra4648/erp/internal/modules/shared/enums"
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
	Items                  []*StockMovementItem       `gorm:"foreignKey:StockMovementID"`
}
