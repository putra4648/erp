package domain

import (
	"putra4648/erp/internal/shared/enums"

	"github.com/google/uuid"
)

type ApprovalTransaction struct {
	ID          uuid.UUID            `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	WorkflowID  uuid.UUID            `gorm:"type:uuid;not null"`
	ReferenceID uuid.UUID            `gorm:"type:uuid;not null;index"` // ID StockAdjustment/PO
	CurrentStep int                  `gorm:"default:1"`
	Status      enums.ApprovalStatus `gorm:"default:'PENDING';size:20"` // PENDING, APPROVED, REJECTED
}
