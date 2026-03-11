package repository

import (
	"context"
	sharedDto "putra4648/erp/internal/shared/dto"
	"putra4648/erp/internal/stock_level/domain"
	"putra4648/erp/internal/stock_level/dto"

	"github.com/google/uuid"
)

type StockLevelRepository interface {
	UpdateQuantity(ctx context.Context, dto *dto.StockLevelDto) error
	Create(ctx context.Context, dto *domain.StockLevel) error

	FindByID(ctx context.Context, id uuid.UUID) (*domain.StockLevel, error)
	FindByProductAndWarehouse(ctx context.Context, dto *dto.StockLevelDto) (*domain.StockLevel, error)
	FindByProductAndWarehouseWithPreload(ctx context.Context, dto *dto.StockLevelDto) (*domain.StockLevel, error)
	FindStockLevels(ctx context.Context, pagination *sharedDto.PaginationRequest, dto *dto.StockLevelDto) ([]*domain.StockLevel, int64, error)
}
