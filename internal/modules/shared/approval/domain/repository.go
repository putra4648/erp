package domain

import "github.com/google/uuid"

type ApprovalRepository interface {
	Create(docCode string, referenceID uuid.UUID) (ApprovalTransaction, error)
}
