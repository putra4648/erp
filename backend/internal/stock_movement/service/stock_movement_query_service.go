package service

import (
	"context"
	sharedDto "putra4648/erp/internal/shared/dto"
	"putra4648/erp/internal/stock_movement/dto"
	"putra4648/erp/internal/stock_movement/mapper"
	"putra4648/erp/internal/stock_movement/repository"

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

func (s *stockMovementQueryService) FindAll(ctx context.Context, pagination *sharedDto.PaginationRequest, req *dto.StockMovementDTO) (*sharedDto.PaginationResponse[*dto.StockMovementDTO], error) {
	models, total, err := s.repo.FindAll(ctx, pagination, req)
	if err != nil {
		return nil, err
	}

	return &sharedDto.PaginationResponse[*dto.StockMovementDTO]{
		Items: mapper.ToDTOs(models),
		Total: total,
		Page:  pagination.Page,
		Size:  pagination.Size,
	}, nil
}

func (s *stockMovementQueryService) FindTransactions(ctx context.Context, pagination *sharedDto.PaginationRequest, req *dto.StockTransactionDTO) (*sharedDto.PaginationResponse[*dto.StockTransactionDTO], error) {
	models, total, err := s.repo.FindTransactions(ctx, pagination, req)
	if err != nil {
		return nil, err
	}

	return &sharedDto.PaginationResponse[*dto.StockTransactionDTO]{
		Items: mapper.ToTransactionDTOs(models),
		Total: total,
		Page:  pagination.Page,
		Size:  pagination.Size,
	}, nil
}
