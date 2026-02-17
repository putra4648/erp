package domain

import (
	"context"
	"putra4648/erp/internal/modules/stock_movement/dto"

	"github.com/google/uuid"
)

type StockMovementRepository interface {
	Create(ctx context.Context, movement *StockMovement) error
	FindByID(ctx context.Context, id uuid.UUID) (*StockMovement, error)
	FindAll(ctx context.Context, req *dto.StockMovementRequest) ([]*StockMovement, int64, error)
	Update(ctx context.Context, movement *StockMovement) error
	Delete(ctx context.Context, id uuid.UUID) error
}
