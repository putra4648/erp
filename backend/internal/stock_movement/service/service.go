package service

import (
	"context"
	sharedDto "putra4648/erp/internal/shared/dto"
	"putra4648/erp/internal/stock_movement/dto"

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
	FindAll(ctx context.Context, pagination *sharedDto.PaginationRequest, req *dto.StockMovementDTO) (*sharedDto.PaginationResponse[*dto.StockMovementDTO], error)
	FindTransactions(ctx context.Context, pagination *sharedDto.PaginationRequest, req *dto.StockTransactionDTO) (*sharedDto.PaginationResponse[*dto.StockTransactionDTO], error)
}
