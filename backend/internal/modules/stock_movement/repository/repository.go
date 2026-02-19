package repository

import (
	"context"
	"putra4648/erp/internal/modules/stock_movement/domain"
	"putra4648/erp/internal/modules/stock_movement/dto"

	"github.com/google/uuid"
)

type StockMovementRepository interface {
	Create(ctx context.Context, movement *domain.StockMovement) error
	FindByID(ctx context.Context, id uuid.UUID) (*domain.StockMovement, error)
	FindAll(ctx context.Context, req *dto.StockMovementRequest) ([]*domain.StockMovement, int64, error)
	Update(ctx context.Context, movement *domain.StockMovement) error
	Delete(ctx context.Context, id uuid.UUID) error
	CompletedMovement(ctx context.Context, id uuid.UUID) error
	CreateTransaction(ctx context.Context, transaction *domain.StockTransaction) error
	FindTransactions(ctx context.Context, req *dto.StockTransactionRequest) ([]*domain.StockTransaction, int64, error)
}
