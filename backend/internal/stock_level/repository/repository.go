package repository

import (
	"context"
	"putra4648/erp/internal/stock_level/domain"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type StockLevelRepository interface {
	GetByProductAndWarehouse(ctx context.Context, productID, warehouseID uuid.UUID) (*domain.StockLevel, error)
	GetByProductAndWarehouseWithPreload(ctx context.Context, productID, warehouseID uuid.UUID) (*domain.StockLevel, error)
	UpdateQuantity(ctx context.Context, productID, warehouseID uuid.UUID, quantity decimal.Decimal) error
	GetStockLevels(ctx context.Context, warehouseID *uuid.UUID, productID *uuid.UUID, search string, page, size int) ([]*domain.StockLevel, int64, error)
	FindByID(ctx context.Context, id uuid.UUID) (*domain.StockLevel, error)
	Save(ctx context.Context, stockLevel *domain.StockLevel) error
}
