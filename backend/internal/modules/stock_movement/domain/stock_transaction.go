package domain

import (
	"time"

	productDomain "putra4648/erp/internal/modules/product/domain"
	supplierDomain "putra4648/erp/internal/modules/supplier/domain"
	warehouseDomain "putra4648/erp/internal/modules/warehouse/domain"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type StockTransaction struct {
	ID           uuid.UUID                 `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ProductID    uuid.UUID                 `gorm:"type:uuid;not null"`
	Product      productDomain.Product     `gorm:"foreignKey:ProductID"`
	WarehouseID  uuid.UUID                 `gorm:"type:uuid;not null"`
	Warehouse    warehouseDomain.Warehouse `gorm:"foreignKey:WarehouseID"`
	SupplierID   uuid.UUID                 `gorm:"type:uuid;not null"`
	Supplier     supplierDomain.Supplier   `gorm:"foreignKey:SupplierID"`
	Type         string                    `gorm:"type:varchar(20)"` // IN, OUT, ADJUST
	Quantity     decimal.Decimal           `gorm:"type:decimal(19,4);not null"`
	BalanceAfter decimal.Decimal           `gorm:"->;column:balance_after"`
	ReferenceNo  string                    `gorm:"type:varchar(255)"`
	CreatedAt    time.Time                 `gorm:"autoCreateTime"`
}
