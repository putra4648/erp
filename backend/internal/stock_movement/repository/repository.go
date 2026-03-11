package repository

import (
	"context"
	sharedDto "putra4648/erp/internal/shared/dto"
	"putra4648/erp/internal/stock_movement/domain"
	"putra4648/erp/internal/stock_movement/dto"

	"github.com/google/uuid"
)

type StockMovementRepository interface {
	Create(ctx context.Context, movement *domain.StockMovement) error
	Update(ctx context.Context, movement *domain.StockMovement) error
	Delete(ctx context.Context, id uuid.UUID) error
	CompletedMovement(ctx context.Context, id uuid.UUID) error
	CreateTransaction(ctx context.Context, transaction *domain.StockTransaction) error

	FindByID(ctx context.Context, id uuid.UUID) (*domain.StockMovement, error)
	FindAll(ctx context.Context, pagination *sharedDto.PaginationRequest, req *dto.StockMovementDTO) ([]*domain.StockMovement, int64, error)
	FindTransactions(ctx context.Context, pagination *sharedDto.PaginationRequest, req *dto.StockTransactionDTO) ([]*domain.StockTransaction, int64, error)
}
