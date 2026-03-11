package service

import (
	"context"
	"putra4648/erp/internal/stock_level/dto"
	"putra4648/erp/internal/stock_level/repository"
)

type stockLevelCommandService struct {
	repo repository.StockLevelRepository
}

func NewStockLevelCommandService(repo repository.StockLevelRepository) StockLevelCommandService {
	return &stockLevelCommandService{
		repo: repo,
	}
}

func (s *stockLevelCommandService) AdjustStock(ctx context.Context, dto *dto.StockLevelDto) error {
	stock, err := s.repo.FindByProductAndWarehouse(ctx, dto)
	if err != nil {
		return err
	}

	newQuantity := dto.Quantity
	if stock != nil {
		newQuantity = stock.Quantity.Add(dto.Quantity)
		dto.Quantity = newQuantity
	}

	return s.repo.UpdateQuantity(ctx, dto)
}
