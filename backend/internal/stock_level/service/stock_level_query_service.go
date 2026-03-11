package service

import (
	"context"
	sharedDto "putra4648/erp/internal/shared/dto"
	"putra4648/erp/internal/stock_level/domain"
	"putra4648/erp/internal/stock_level/dto"
	"putra4648/erp/internal/stock_level/repository"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type stockLevelQueryService struct {
	repo repository.StockLevelRepository
}

func NewStockLevelQueryService(repo repository.StockLevelRepository) StockLevelQueryService {
	return &stockLevelQueryService{
		repo: repo,
	}
}

func (s *stockLevelQueryService) FindStockLevelQuantity(ctx context.Context, dto *dto.StockLevelDto) (decimal.Decimal, error) {
	stock, err := s.repo.FindByProductAndWarehouse(ctx, dto)
	if err != nil {
		return decimal.Zero, err
	}
	if stock == nil {
		return decimal.Zero, nil
	}
	return stock.Quantity, nil
}

func (s *stockLevelQueryService) FindAllStockLevels(ctx context.Context, pagination *sharedDto.PaginationRequest, dto *dto.StockLevelDto) (*sharedDto.PaginationResponse[*domain.StockLevel], error) {
	stocks, total, err := s.repo.FindStockLevels(ctx, pagination, dto)
	if err != nil {
		return nil, err
	}

	var responses []*domain.StockLevel
	for _, stock := range stocks {
		responses = append(responses, stock)
	}
	return &sharedDto.PaginationResponse[*domain.StockLevel]{
		Items: responses,
		Total: total,
		Page:  pagination.Page,
		Size:  pagination.Size,
	}, nil
}

func (s *stockLevelQueryService) FindByProductAndWarehouseWithPreload(ctx context.Context, dto *dto.StockLevelDto) (*domain.StockLevel, error) {
	return s.repo.FindByProductAndWarehouseWithPreload(ctx, dto)
}

func (s *stockLevelQueryService) FindByID(ctx context.Context, id uuid.UUID) (*domain.StockLevel, error) {
	return s.repo.FindByID(ctx, id)
}
