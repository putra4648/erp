package repository

import (
	"context"

	"github.com/google/uuid"

	"putra4648/erp/internal/modules/stock_adjustment/domain"
)

type StockAdjustmentRepository interface {
	Save(ctx context.Context, adjustment *domain.StockAdjustment) error
	FindByID(ctx context.Context, id uuid.UUID) (*domain.StockAdjustment, error)
	FindAll(ctx context.Context, page, size int) ([]*domain.StockAdjustment, int64, error)
	Update(ctx context.Context, adjustment *domain.StockAdjustment) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type AdjustmentReasonRepository interface {
	Save(ctx context.Context, reason *domain.AdjustmentReason) error
	FindByID(ctx context.Context, id uuid.UUID) (*domain.AdjustmentReason, error)
	FindAll(ctx context.Context) ([]*domain.AdjustmentReason, error)
	Update(ctx context.Context, reason *domain.AdjustmentReason) error
	Delete(ctx context.Context, id uuid.UUID) error
}
