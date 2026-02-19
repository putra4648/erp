package service

import (
	"context"
	"putra4648/erp/internal/modules/stock_movement/dto"

	"github.com/google/uuid"
)

type StockMovementCommandService interface {
	Create(ctx context.Context, dto *dto.StockMovementDTO) (*dto.StockMovementDTO, error)
	Update(ctx context.Context, id uuid.UUID, dto *dto.StockMovementDTO) (*dto.StockMovementDTO, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Approve(ctx context.Context, id uuid.UUID) error
}

type StockMovementQueryService interface {
	FindByID(ctx context.Context, id uuid.UUID) (*dto.StockMovementDTO, error)
	FindAll(ctx context.Context, req *dto.StockMovementRequest) ([]*dto.StockMovementDTO, int64, error)
	FindTransactions(ctx context.Context, req *dto.StockTransactionRequest) ([]*dto.StockTransactionResponse, int64, error)
}
