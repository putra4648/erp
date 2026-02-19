package service

import (
	"context"
	"putra4648/erp/internal/modules/stock_adjustment/dto"
	"putra4648/erp/internal/modules/stock_adjustment/mapper"
	"putra4648/erp/internal/modules/stock_adjustment/repository"
)

type adjustmentReasonQueryService struct {
	repo repository.AdjustmentReasonRepository
}

func NewAdjustmentReasonQueryService(repo repository.AdjustmentReasonRepository) AdjustmentReasonQueryService {
	return &adjustmentReasonQueryService{repo: repo}
}

func (s *adjustmentReasonQueryService) FindAll(ctx context.Context) ([]*dto.AdjustmentReasonDto, error) {
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
