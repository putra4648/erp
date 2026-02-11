package domain

import "github.com/google/uuid"

type ApprovalStep struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	WorkflowID      uuid.UUID `gorm:"type:uuid;not null"`
	StepOrder       int       `gorm:"not null"`
	TargetGroupName string    `gorm:"not null;size:100"` // Dari Keycloak
	MinApprover     int       `gorm:"default:1"`
}
