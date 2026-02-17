package domain

import (
	"time"

	"github.com/google/uuid"
)

type StockLevel struct {
	ID          uuid.UUID `gorm:"type:uuid;primarykey;default:uuid_generate_v4()"`
	ProductID   uuid.UUID `gorm:"type:uuid;not null"`
	WarehouseID uuid.UUID `gorm:"type:uuid;not null"`
	Quantity    int       `gorm:"type:int;not null"`
	LastUpdated time.Time `gorm:"type:timestamp;not null"`
}
