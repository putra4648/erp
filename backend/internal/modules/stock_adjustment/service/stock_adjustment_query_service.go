package service

import (
	"context"
	"putra4648/erp/internal/modules/stock_adjustment/dto"
	"putra4648/erp/internal/modules/stock_adjustment/mapper"
	"putra4648/erp/internal/modules/stock_adjustment/repository"

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

func (s *stockAdjustmentQueryService) FindAll(ctx context.Context, page, size int) ([]*dto.StockAdjustmentDto, int64, error) {
	adjustments, total, err := s.repo.FindAll(ctx, page, size)
	if err != nil {
		return nil, 0, err
	}

	var responses []*dto.StockAdjustmentDto
	for _, adj := range adjustments {
		responses = append(responses, mapper.ToStockAdjustmentDto(adj))
	}
	return responses, total, nil
}
