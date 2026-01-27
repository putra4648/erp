package repositories

import (
	"putra4648/erp/internal/models"

	"gorm.io/gorm"
)

type inventoryRepository struct {
	db *gorm.DB
}

type InventoryRepository interface {
	CreateStockAdjustment(stockAdjustment models.StockAdjustment) error
}

func NewInventoryRepository(db *gorm.DB) *inventoryRepository {
	return &inventoryRepository{db: db}
}

func (r *inventoryRepository) CreateStockAdjustment(stockAdjustment models.StockAdjustment) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		return nil
	})
}
