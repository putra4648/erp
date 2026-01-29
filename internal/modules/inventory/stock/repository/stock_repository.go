package repository

import (
	"putra4648/erp/internal/modules/inventory/stock/domain"
	"putra4648/erp/internal/modules/inventory/stock/dto"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type stockRepo struct {
	db *gorm.DB
}

// NewStockRepository mengembalikan interface domain.StockRepository
// Ini memastikan layer service hanya tahu interface-nya saja
func NewStockRepository(db *gorm.DB) *stockRepo {
	return &stockRepo{db}
}

func (r *stockRepo) CreateStockAdjustment(dto dto.StockRequest, userID uuid.UUID) (domain.StockAdjustment, error) {
	var res domain.StockAdjustment
	err := r.db.Transaction(func(tx *gorm.DB) error {
		adjustment := domain.StockAdjustment{
			ID:        uuid.New(),
			ProductID: uuid.MustParse(dto.ProductID),
			QtyDiff:   decimal.NewFromFloat(dto.Qty),
			Reason:    dto.Reason,
			Status:    "WAITING_APPROVAL",
			CreatedBy: userID,
		}

		if err := tx.Create(&adjustment).Error; err != nil {
			return err
		}

		res = adjustment

		return nil
	})

	return res, err
}
