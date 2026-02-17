package domain

import (
	"github.com/google/uuid"
)

type AdjustmentReason struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	AccountCode string    `gorm:"type:varchar(20);not null" json:"account_code"`
}
