package service

import (
	"context"
	sharedDto "putra4648/erp/internal/shared/dto"
	"putra4648/erp/internal/stock_level/domain"
	"putra4648/erp/internal/stock_level/dto"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type StockLevelQueryService interface {
	FindStockLevelQuantity(ctx context.Context, dto *dto.StockLevelDto) (decimal.Decimal, error)
	FindAllStockLevels(ctx context.Context, pagination *sharedDto.PaginationRequest, dto *dto.StockLevelDto) (*sharedDto.PaginationResponse[*domain.StockLevel], error)
	FindByProductAndWarehouseWithPreload(ctx context.Context, dto *dto.StockLevelDto) (*domain.StockLevel, error)
	FindByID(ctx context.Context, id uuid.UUID) (*domain.StockLevel, error)
}

type StockLevelCommandService interface {
	AdjustStock(ctx context.Context, dto *dto.StockLevelDto) error
}
