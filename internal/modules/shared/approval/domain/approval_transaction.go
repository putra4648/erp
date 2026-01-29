package domain

import "github.com/google/uuid"

type ApprovalTransaction struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	WorkflowID  uuid.UUID `gorm:"type:uuid;not null"`
	ReferenceID uuid.UUID `gorm:"type:uuid;not null;index"` // ID StockAdjustment/PO
	CurrentStep int       `gorm:"default:1"`
	Status      string    `gorm:"default:'PENDING';size:20"` // PENDING, APPROVED, REJECTED
}
