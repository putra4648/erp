package service

import (
	"context"
	"putra4648/erp/internal/stock_level/repository"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type stockLevelCommandService struct {
	repo repository.StockLevelRepository
}

func NewStockLevelCommandService(repo repository.StockLevelRepository) StockLevelCommandService {
	return &stockLevelCommandService{
		repo: repo,
	}
}

func (s *stockLevelCommandService) AdjustStock(ctx context.Context, productID, warehouseID uuid.UUID, delta decimal.Decimal) error {
	stock, err := s.repo.GetByProductAndWarehouse(ctx, productID, warehouseID)
	if err != nil {
		return err
	}

	newQuantity := delta
	if stock != nil {
		newQuantity = stock.Quantity.Add(delta)
	}

	return s.repo.UpdateQuantity(ctx, productID, warehouseID, newQuantity)
}
