package service

import (
	"context"
	"putra4648/erp/internal/modules/stock_adjustment/domain"
	"putra4648/erp/internal/modules/stock_adjustment/dto"

	"github.com/google/uuid"
)

type AdjustmentReasonService interface {
	Create(ctx context.Context, req *dto.AdjustmentReasonRequest) (*domain.AdjustmentReason, error)
	FindAll(ctx context.Context) ([]*domain.AdjustmentReason, error)
}

type adjustmentReasonService struct {
	repo domain.AdjustmentReasonRepository
}

func NewAdjustmentReasonService(repo domain.AdjustmentReasonRepository) AdjustmentReasonService {
	return &adjustmentReasonService{repo: repo}
}

func (s *adjustmentReasonService) Create(ctx context.Context, req *dto.AdjustmentReasonRequest) (*domain.AdjustmentReason, error) {
	reason := &domain.AdjustmentReason{
		ID:          uuid.New(),
		Name:        req.Name,
		AccountCode: req.AccountCode,
	}
	if err := s.repo.Save(ctx, reason); err != nil {
		return nil, err
	}
	return reason, nil
}

func (s *adjustmentReasonService) FindAll(ctx context.Context) ([]*domain.AdjustmentReason, error) {
	return s.repo.FindAll(ctx)
}
