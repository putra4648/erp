package models

import (
	"time"

	"github.com/google/uuid"
)

type ApprovalWorkflow struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	DocCode   string         `gorm:"unique;not null;size:50"` // Contoh: 'STOCK_ADJ'
	DocName   string         `gorm:"not null;size:100"`
	IsActive  bool           `gorm:"default:true"`
	Steps     []ApprovalStep `gorm:"foreignKey:WorkflowID"`
	CreatedAt time.Time
}
