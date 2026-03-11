package service

import (
	"context"
	"putra4648/erp/internal/stock_adjustment/dto"
	"putra4648/erp/internal/stock_adjustment/mapper"
	"putra4648/erp/internal/stock_adjustment/repository"

	sharedDto "putra4648/erp/internal/shared/dto"
)

type adjustmentReasonQueryService struct {
	repo repository.AdjustmentReasonRepository
}

func NewAdjustmentReasonQueryService(repo repository.AdjustmentReasonRepository) AdjustmentReasonQueryService {
	return &adjustmentReasonQueryService{repo: repo}
}

func (s *adjustmentReasonQueryService) FindAll(ctx context.Context, pagination *sharedDto.PaginationRequest, req *dto.AdjustmentReasonDto) (*sharedDto.PaginationResponse[*dto.AdjustmentReasonDto], error) {
	reasons, total, err := s.repo.FindAll(ctx, pagination, req)
	if err != nil {
		return nil, err
	}

	var responses []*dto.AdjustmentReasonDto
	for _, reason := range reasons {
		responses = append(responses, mapper.ToAdjustmentReasonDto(reason))
	}
	return &sharedDto.PaginationResponse[*dto.AdjustmentReasonDto]{
		Items: responses,
		Total: total,
		Page:  pagination.Page,
		Size:  pagination.Size,
	}, nil
}
