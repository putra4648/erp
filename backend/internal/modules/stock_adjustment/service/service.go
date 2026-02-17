package service

import (
	"context"
	"putra4648/erp/internal/modules/stock_adjustment/dto"

	"github.com/google/uuid"
)

type StockAdjustmentService interface {
	Create(ctx context.Context, req *dto.CreateStockAdjustmentRequest, userID uuid.UUID) (*dto.StockAdjustmentDto, error)
	FindByID(ctx context.Context, id uuid.UUID) (*dto.StockAdjustmentDto, error)
	FindAll(ctx context.Context, page, size int) ([]*dto.StockAdjustmentDto, int64, error)
	Update(ctx context.Context, id uuid.UUID, req *dto.CreateStockAdjustmentRequest) (*dto.StockAdjustmentDto, error)
	Approve(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*dto.StockAdjustmentDto, error)
	Void(ctx context.Context, id uuid.UUID) (*dto.StockAdjustmentDto, error)
}

type AdjustmentReasonService interface {
	Create(ctx context.Context, req *dto.AdjustmentReasonRequest) (*dto.AdjustmentReasonDto, error)
	FindAll(ctx context.Context) ([]*dto.AdjustmentReasonDto, error)
}
