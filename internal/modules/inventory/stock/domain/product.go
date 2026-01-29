package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Product struct {
	ID        uuid.UUID       `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	SKU       string          `gorm:"unique;not null;size:50"`
	Name      string          `gorm:"not null;size:255"`
	UOM       string          `gorm:"not null;size:20"` // Unit of Measure
	Price     decimal.Decimal `gorm:"type:decimal(19,4)"`
	CreatedAt time.Time
}
