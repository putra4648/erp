package service

import (
	"putra4648/erp/internal/modules/shared/approval/domain"

	"github.com/google/uuid"
	"go.uber.org/dig"
)

type ApprovalService struct {
	repository domain.ApprovalRepository
}

type ApprovalServiceParams struct {
	dig.In
	Repository domain.ApprovalRepository
}

func NewApprovalService(params ApprovalServiceParams) *ApprovalService {
	return &ApprovalService{
		repository: params.Repository,
	}
}

func (s *ApprovalService) SubmitApproval(docCode string, referenceID uuid.UUID) (domain.ApprovalTransaction, error) {
	return s.repository.Create(docCode, referenceID)
}
