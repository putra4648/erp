package service

import (
	"context"
	"putra4648/erp/internal/modules/stock_level/domain"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type StockLevelQueryService interface {
	GetStockLevel(ctx context.Context, productID, warehouseID uuid.UUID) (decimal.Decimal, error)
	GetAllStockLevels(ctx context.Context, warehouseID *uuid.UUID, productID *uuid.UUID, search string, page, size int) ([]*domain.StockLevel, int64, error)
	GetByProductAndWarehouseWithPreload(ctx context.Context, productID, warehouseID uuid.UUID) (*domain.StockLevel, error)
	FindByID(ctx context.Context, id uuid.UUID) (*domain.StockLevel, error)
}

type StockLevelCommandService interface {
	AdjustStock(ctx context.Context, productID, warehouseID uuid.UUID, delta decimal.Decimal) error
}
