package service

import (
	"context"
	"putra4648/erp/internal/modules/stock_movement/dto"
	"putra4648/erp/internal/modules/stock_movement/mapper"
	"putra4648/erp/internal/modules/stock_movement/repository"

	"github.com/google/uuid"
)

type stockMovementQueryService struct {
	repo repository.StockMovementRepository
}

func NewStockMovementQueryService(repo repository.StockMovementRepository) StockMovementQueryService {
	return &stockMovementQueryService{repo: repo}
}

func (s *stockMovementQueryService) FindByID(ctx context.Context, id uuid.UUID) (*dto.StockMovementDTO, error) {
	model, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapper.ToDTO(model), nil
}

func (s *stockMovementQueryService) FindAll(ctx context.Context, req *dto.StockMovementRequest) ([]*dto.StockMovementDTO, int64, error) {
	models, total, err := s.repo.FindAll(ctx, req)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]*dto.StockMovementDTO, len(models))
	for i, m := range models {
		responses[i] = mapper.ToDTO(m)
	}

	return responses, total, nil
}

func (s *stockMovementQueryService) FindTransactions(ctx context.Context, req *dto.StockTransactionRequest) ([]*dto.StockTransactionResponse, int64, error) {
	models, total, err := s.repo.FindTransactions(ctx, req)
	if err != nil {
		return nil, 0, err
	}

	return mapper.ToTransactionDTOs(models), total, nil
}
