package service

import (
	"context"
	"putra4648/erp/internal/stock_adjustment/domain"
	"putra4648/erp/internal/stock_adjustment/dto"
	"putra4648/erp/internal/stock_adjustment/mapper"
	"putra4648/erp/internal/stock_adjustment/repository"

	"github.com/google/uuid"
)

type adjustmentReasonCommandService struct {
	repo repository.AdjustmentReasonRepository
}

func NewAdjustmentReasonCommandService(repo repository.AdjustmentReasonRepository) AdjustmentReasonCommandService {
	return &adjustmentReasonCommandService{repo: repo}
}

func (s *adjustmentReasonCommandService) Create(ctx context.Context, req *dto.AdjustmentReasonRequest) (*dto.AdjustmentReasonDto, error) {
	reason := &domain.AdjustmentReason{
		ID:   uuid.New(),
		Name: req.Name,
	}

	if err := s.repo.Save(ctx, reason); err != nil {
		return nil, err
	}

	return mapper.ToAdjustmentReasonDto(reason), nil
}
