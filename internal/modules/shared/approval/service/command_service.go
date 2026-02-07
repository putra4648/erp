package service

import (
	"context"
	"putra4648/erp/internal/modules/shared/approval/domain"
	"putra4648/erp/internal/modules/shared/approval/repository"

	"github.com/google/uuid"
	"go.uber.org/dig"
)

type ApprovalService struct {
	repository repository.ApprovalRepository
}

type ApprovalServiceParams struct {
	dig.In
	Repository repository.ApprovalRepository
}

func NewApprovalService(params ApprovalServiceParams) *ApprovalService {
	return &ApprovalService{
		repository: params.Repository,
	}
}

func (s *ApprovalService) SubmitApproval(ctx context.Context, docCode string, referenceID uuid.UUID) (domain.ApprovalTransaction, error) {
	return s.repository.Create(ctx, docCode, referenceID)
}
