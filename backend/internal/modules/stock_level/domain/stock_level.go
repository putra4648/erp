package domain

import (
	"time"

	productDomain "putra4648/erp/internal/modules/product/domain"
	warehouseDomain "putra4648/erp/internal/modules/warehouse/domain"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type StockLevel struct {
	ID          uuid.UUID                 `gorm:"type:uuid;primarykey;default:uuid_generate_v4()"`
	ProductID   uuid.UUID                 `gorm:"type:uuid;not null"`
	Product     productDomain.Product     `gorm:"foreignKey:ProductID"`
	WarehouseID uuid.UUID                 `gorm:"type:uuid;not null"`
	Warehouse   warehouseDomain.Warehouse `gorm:"foreignKey:WarehouseID"`
	Quantity    decimal.Decimal           `gorm:"type:decimal(19,4);not null"`
	LastUpdated time.Time                 `gorm:"type:timestamp;not null"`
}
