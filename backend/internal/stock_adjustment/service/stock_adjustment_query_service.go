package service

import (
	"context"
	sharedDto "putra4648/erp/internal/shared/dto"
	"putra4648/erp/internal/stock_adjustment/dto"
	"putra4648/erp/internal/stock_adjustment/mapper"
	"putra4648/erp/internal/stock_adjustment/repository"

	"github.com/google/uuid"
)

type stockAdjustmentQueryService struct {
	repo repository.StockAdjustmentRepository
}

func NewStockAdjustmentQueryService(repo repository.StockAdjustmentRepository) StockAdjustmentQueryService {
	return &stockAdjustmentQueryService{repo: repo}
}

func (s *stockAdjustmentQueryService) FindByID(ctx context.Context, id uuid.UUID) (*dto.StockAdjustmentDto, error) {
	adjustment, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapper.ToStockAdjustmentDto(adjustment), nil
}

func (s *stockAdjustmentQueryService) FindAll(ctx context.Context, pagination *sharedDto.PaginationRequest, req *dto.StockAdjustmentDto) (*sharedDto.PaginationResponse[*dto.StockAdjustmentDto], error) {
	adjustments, total, err := s.repo.FindAll(ctx, pagination, req)
	if err != nil {
		return nil, err
	}

	var responses []*dto.StockAdjustmentDto
	for _, adj := range adjustments {
		responses = append(responses, mapper.ToStockAdjustmentDto(adj))
	}
	return &sharedDto.PaginationResponse[*dto.StockAdjustmentDto]{
		Items: responses,
		Total: total,
		Page:  pagination.Page,
		Size:  pagination.Size,
	}, nil
}
