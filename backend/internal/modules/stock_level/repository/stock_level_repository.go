package repository

import (
	"context"
	"putra4648/erp/internal/modules/stock_level/domain"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type stockLevelRepository struct {
	db *gorm.DB
}

func NewStockLevelRepository(db *gorm.DB) StockLevelRepository {
	return &stockLevelRepository{db: db}
}

func (r *stockLevelRepository) GetByProductAndWarehouse(ctx context.Context, productID, warehouseID uuid.UUID) (*domain.StockLevel, error) {
	var stockLevel domain.StockLevel
	err := r.db.WithContext(ctx).
		Where("product_id = ? AND warehouse_id = ?", productID, warehouseID).
		First(&stockLevel).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &stockLevel, nil
}

func (r *stockLevelRepository) GetByProductAndWarehouseWithPreload(ctx context.Context, productID, warehouseID uuid.UUID) (*domain.StockLevel, error) {
	var stockLevel domain.StockLevel
	err := r.db.WithContext(ctx).
		Preload("Product").Preload("Warehouse").
		Where("product_id = ? AND warehouse_id = ?", productID, warehouseID).
		First(&stockLevel).Error
	if err != nil {
		return nil, err
	}
	return &stockLevel, nil
}

func (r *stockLevelRepository) UpdateQuantity(ctx context.Context, productID, warehouseID uuid.UUID, quantity decimal.Decimal) error {
	var stockLevel domain.StockLevel
	err := r.db.WithContext(ctx).
		Where("product_id = ? AND warehouse_id = ?", productID, warehouseID).
		First(&stockLevel).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create new stock level
			stockLevel = domain.StockLevel{
				ID:          uuid.New(),
				ProductID:   productID,
				WarehouseID: warehouseID,
				Quantity:    quantity,
				LastUpdated: time.Now(),
			}
			return r.db.WithContext(ctx).Create(&stockLevel).Error
		}
		return err
	}

	// Update existing stock level
	// We use the new quantity (it could be absolute or relative depending on how we call this,
	// but usually derived from movements we want to ATOMICALLY add/subtract).
	// However, for simplicity now let's say this is the new total.
	stockLevel.Quantity = quantity
	stockLevel.LastUpdated = time.Now()
	return r.db.WithContext(ctx).Save(&stockLevel).Error
}

func (r *stockLevelRepository) GetStockLevels(ctx context.Context, warehouseID *uuid.UUID, productID *uuid.UUID, search string, page, size int) ([]*domain.StockLevel, int64, error) {
	var stockLevels []*domain.StockLevel
	var total int64
	query := r.db.WithContext(ctx).Model(&domain.StockLevel{}).Preload("Product").Preload("Warehouse")

	if warehouseID != nil {
		query = query.Where("warehouse_id = ?", *warehouseID)
	}
	if productID != nil {
		query = query.Where("product_id = ?", *productID)
	}

	if search != "" {
		query = query.Joins("JOIN products ON products.id = stock_levels.product_id").
			Where("products.name ILIKE ?", "%"+search+"%")
	}

	query.Count(&total)

	if page > 0 && size > 0 {
		offset := (page - 1) * size
		query = query.Limit(size).Offset(offset)
	}

	err := query.Find(&stockLevels).Error
	return stockLevels, total, err
}

func (r *stockLevelRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.StockLevel, error) {
	var stockLevel domain.StockLevel
	err := r.db.WithContext(ctx).Preload("Product").Preload("Warehouse").Where("id = ?", id).First(&stockLevel).Error
	if err != nil {
		return nil, err
	}
	return &stockLevel, nil
}

func (r *stockLevelRepository) Save(ctx context.Context, stockLevel *domain.StockLevel) error {
	if stockLevel.ID == uuid.Nil {
		stockLevel.ID = uuid.New()
	}
	stockLevel.LastUpdated = time.Now()
	return r.db.WithContext(ctx).Save(stockLevel).Error
}
