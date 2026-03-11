package repository

import (
	"context"
	sharedDto "putra4648/erp/internal/shared/dto"
	"putra4648/erp/internal/stock_adjustment/domain"
	"putra4648/erp/internal/stock_adjustment/dto"

	"github.com/google/uuid"
)

type StockAdjustmentRepository interface {
	Save(ctx context.Context, adjustment *domain.StockAdjustment) error
	Update(ctx context.Context, adjustment *domain.StockAdjustment) error
	Delete(ctx context.Context, id uuid.UUID) error

	FindByID(ctx context.Context, id uuid.UUID) (*domain.StockAdjustment, error)
	FindAll(ctx context.Context, pagination *sharedDto.PaginationRequest, dto *dto.StockAdjustmentDto) ([]*domain.StockAdjustment, int64, error)
}

type AdjustmentReasonRepository interface {
	Save(ctx context.Context, reason *domain.AdjustmentReason) error
	Update(ctx context.Context, reason *domain.AdjustmentReason) error
	Delete(ctx context.Context, id uuid.UUID) error

	FindByID(ctx context.Context, id uuid.UUID) (*domain.AdjustmentReason, error)
	FindAll(ctx context.Context, pagination *sharedDto.PaginationRequest, dto *dto.AdjustmentReasonDto) ([]*domain.AdjustmentReason, int64, error)
}
