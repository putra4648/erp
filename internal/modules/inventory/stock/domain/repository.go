package domain

import (
	"putra4648/erp/internal/modules/inventory/stock/dto"

	"github.com/google/uuid"
)

type StockRepository interface {
	CreateStockAdjustment(dto dto.StockRequest, userID uuid.UUID) (StockAdjustment, error)
}
