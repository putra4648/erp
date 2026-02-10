package domain

import (
	"context"

	"github.com/google/uuid"
)

type StockAdjustmentRepository interface {
	Save(ctx context.Context, adjustment *StockAdjustment) error
	FindByID(ctx context.Context, id uuid.UUID) (*StockAdjustment, error)
	FindAll(ctx context.Context, page, size int) ([]*StockAdjustment, int64, error)
	Update(ctx context.Context, adjustment *StockAdjustment) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type AdjustmentReasonRepository interface {
	Save(ctx context.Context, reason *AdjustmentReason) error
	FindByID(ctx context.Context, id uuid.UUID) (*AdjustmentReason, error)
	FindAll(ctx context.Context) ([]*AdjustmentReason, error)
	Update(ctx context.Context, reason *AdjustmentReason) error
	Delete(ctx context.Context, id uuid.UUID) error
}
