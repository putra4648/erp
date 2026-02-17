package domain

import (
	productDomain "putra4648/erp/internal/modules/product/domain"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type StockMovementItem struct {
	ID              uuid.UUID             `gorm:"type:uuid;primarykey;default:uuid_generate_v4()"`
	StockMovementID uuid.UUID             `gorm:"type:uuid;column:stock_movement_id;not null"`
	ProductID       uuid.UUID             `gorm:"type:uuid;column:product_id;not null"`
	Product         productDomain.Product `gorm:"foreignKey:ProductID"`
	Quantity        decimal.Decimal       `gorm:"type:decimal(19,4);not null"`
	Note            string                `gorm:"type:text"`
}
