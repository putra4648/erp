package domain

import (
	"context"
	"putra4648/erp/internal/modules/stock_level/dto"

	"github.com/google/uuid"
)

type StockLevelRepository interface {
	Save(ctx context.Context, stockLevel *dto.StockLevelRequest) error
	FindByID(ctx context.Context, id uuid.UUID) (*StockLevel, error)
	FindAll(ctx context.Context, page, size int) ([]*StockLevel, int64, error)
	Update(ctx context.Context, stockLevel *dto.StockLevelRequest) error
	Delete(ctx context.Context, id uuid.UUID) error
}
