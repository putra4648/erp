package service

import (
	"context"
	"putra4648/erp/internal/modules/stock_adjustment/domain"
	"putra4648/erp/internal/modules/stock_adjustment/dto"
	"putra4648/erp/internal/modules/stock_adjustment/mapper"
	"putra4648/erp/internal/modules/stock_adjustment/repository"

	"github.com/google/uuid"
)

type adjustmentReasonService struct {
	repo repository.AdjustmentReasonRepository
}

func NewAdjustmentReasonService(repo repository.AdjustmentReasonRepository) AdjustmentReasonService {
	return &adjustmentReasonService{repo: repo}
}

func (s *adjustmentReasonService) Create(ctx context.Context, req *dto.AdjustmentReasonRequest) (*dto.AdjustmentReasonDto, error) {
	reason := &domain.AdjustmentReason{
		ID:          uuid.New(),
		Name:        req.Name,
		AccountCode: req.AccountCode,
	}
	if err := s.repo.Save(ctx, reason); err != nil {
		return nil, err
	}
	return mapper.ToAdjustmentReasonDto(reason), nil
}

func (s *adjustmentReasonService) FindAll(ctx context.Context) ([]*dto.AdjustmentReasonDto, error) {
	reasons, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var responses []*dto.AdjustmentReasonDto
	for _, reason := range reasons {
		responses = append(responses, mapper.ToAdjustmentReasonDto(reason))
	}
	return responses, nil
}
