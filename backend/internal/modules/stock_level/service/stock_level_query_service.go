package service

import (
	"context"
	"putra4648/erp/internal/modules/stock_level/domain"
	"putra4648/erp/internal/modules/stock_level/repository"

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

func (s *stockLevelQueryService) GetStockLevel(ctx context.Context, productID, warehouseID uuid.UUID) (decimal.Decimal, error) {
	stock, err := s.repo.GetByProductAndWarehouse(ctx, productID, warehouseID)
	if err != nil {
		return decimal.Zero, err
	}
	if stock == nil {
		return decimal.Zero, nil
	}
	return stock.Quantity, nil
}

func (s *stockLevelQueryService) GetAllStockLevels(ctx context.Context, warehouseID *uuid.UUID, productID *uuid.UUID, search string, page, size int) ([]*domain.StockLevel, int64, error) {
	return s.repo.GetStockLevels(ctx, warehouseID, productID, search, page, size)
}

func (s *stockLevelQueryService) GetByProductAndWarehouseWithPreload(ctx context.Context, productID, warehouseID uuid.UUID) (*domain.StockLevel, error) {
	return s.repo.GetByProductAndWarehouseWithPreload(ctx, productID, warehouseID)
}

func (s *stockLevelQueryService) FindByID(ctx context.Context, id uuid.UUID) (*domain.StockLevel, error) {
	return s.repo.FindByID(ctx, id)
}
