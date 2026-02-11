package service

import (
	"context"
	"putra4648/erp/internal/modules/stock_movement/domain"
	"putra4648/erp/internal/modules/stock_movement/dto"
	"putra4648/erp/internal/modules/stock_movement/repository"

	"github.com/google/uuid"
)

type StockMovementQueryService interface {
	FindByID(ctx context.Context, id uuid.UUID) (*domain.StockMovementResponse, error)
	FindAll(ctx context.Context, req *dto.StockMovementRequest) ([]*domain.StockMovementResponse, int64, error)
}

type stockMovementQueryService struct {
	repo repository.StockMovementRepository
}

func NewStockMovementQueryService(repo repository.StockMovementRepository) StockMovementQueryService {
	return &stockMovementQueryService{repo: repo}
}

func (s *stockMovementQueryService) FindByID(ctx context.Context, id uuid.UUID) (*domain.StockMovementResponse, error) {
	model, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return model.ToResponse(), nil
}

func (s *stockMovementQueryService) FindAll(ctx context.Context, req *dto.StockMovementRequest) ([]*domain.StockMovementResponse, int64, error) {
	models, total, err := s.repo.FindAll(ctx, req)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]*domain.StockMovementResponse, len(models))
	for i, m := range models {
		responses[i] = m.ToResponse()
	}

	return responses, total, nil
}
