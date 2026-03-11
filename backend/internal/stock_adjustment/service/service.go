package service

import (
	"context"
	sharedDto "putra4648/erp/internal/shared/dto"
	"putra4648/erp/internal/stock_adjustment/dto"

	"github.com/google/uuid"
)

type StockAdjustmentQueryService interface {
	FindByID(ctx context.Context, id uuid.UUID) (*dto.StockAdjustmentDto, error)
	FindAll(ctx context.Context, pagination *sharedDto.PaginationRequest, dto *dto.StockAdjustmentDto) (*sharedDto.PaginationResponse[*dto.StockAdjustmentDto], error)
}

type StockAdjustmentCommandService interface {
	Create(ctx context.Context, req *dto.CreateStockAdjustmentRequest, userID uuid.UUID) (*dto.StockAdjustmentDto, error)
	Update(ctx context.Context, id uuid.UUID, req *dto.CreateStockAdjustmentRequest) (*dto.StockAdjustmentDto, error)
	Approve(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*dto.StockAdjustmentDto, error)
	Void(ctx context.Context, id uuid.UUID) (*dto.StockAdjustmentDto, error)
}

type AdjustmentReasonQueryService interface {
	FindAll(ctx context.Context, pagination *sharedDto.PaginationRequest, dto *dto.AdjustmentReasonDto) (*sharedDto.PaginationResponse[*dto.AdjustmentReasonDto], error)
}

type AdjustmentReasonCommandService interface {
	Create(ctx context.Context, req *dto.AdjustmentReasonRequest) (*dto.AdjustmentReasonDto, error)
}
