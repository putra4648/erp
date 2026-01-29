package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type StockAdjustment struct {
	ID        uuid.UUID       `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	ProductID uuid.UUID       `gorm:"type:uuid;not null"`
	Product   Product         `gorm:"foreignKey:ProductID"`
	QtyDiff   decimal.Decimal `gorm:"type:decimal(19,4)"`
	Reason    string          `gorm:"type:text"`
	Status    string          `gorm:"default:'DRAFT';size:20"` // DRAFT, WAITING, COMPLETED
	CreatedBy uuid.UUID       `gorm:"type:uuid;not null"`      // Sub dari Keycloak
	CreatedAt time.Time
}
