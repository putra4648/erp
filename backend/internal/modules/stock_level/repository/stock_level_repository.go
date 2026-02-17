package repository

import (
	"context"
	"putra4648/erp/internal/modules/stock_level/domain"
	"putra4648/erp/internal/modules/stock_level/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type stockLevelRepository struct {
	db *gorm.DB
}

// Delete implements [domain.StockLevelRepository].
func (r *stockLevelRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&domain.StockLevel{}, "id = ?", id).Error
}

// FindAll implements [domain.StockLevelRepository].
func (r *stockLevelRepository) FindAll(ctx context.Context, page int, size int) ([]*domain.StockLevel, int64, error) {
	var stockLevels []*domain.StockLevel
	var count int64
	err := r.db.WithContext(ctx).Count(&count).Find(&stockLevels).Error
	if err != nil {
		return nil, 0, err
	}
	return stockLevels, count, nil
}

// FindByID implements [domain.StockLevelRepository].
func (r *stockLevelRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.StockLevel, error) {
	var stockLevel domain.StockLevel
	err := r.db.WithContext(ctx).First(&stockLevel, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &stockLevel, nil
}

// Update implements [domain.StockLevelRepository].
func (r *stockLevelRepository) Update(ctx context.Context, stockLevel *dto.StockLevelRequest) error {
	productID, err := uuid.Parse(stockLevel.ProductID)
	if err != nil {
		return err
	}
	warehouseID, err := uuid.Parse(stockLevel.WarehouseID)
	if err != nil {
		return err
	}
	stockLevelDomain := &domain.StockLevel{
		ProductID:   productID,
		WarehouseID: warehouseID,
		Quantity:    stockLevel.Quantity,
	}
	return r.db.WithContext(ctx).Save(stockLevelDomain).Error
}

func NewStockLevelRepository(db *gorm.DB) domain.StockLevelRepository {
	return &stockLevelRepository{db: db}
}

func (r *stockLevelRepository) Save(ctx context.Context, stockLevel *dto.StockLevelRequest) error {

	productID, err := uuid.Parse(stockLevel.ProductID)
	if err != nil {
		return err
	}
	warehouseID, err := uuid.Parse(stockLevel.WarehouseID)
	if err != nil {
		return err
	}
	stockLevelDomain := &domain.StockLevel{
		ProductID:   productID,
		WarehouseID: warehouseID,
		Quantity:    stockLevel.Quantity,
	}
	return r.db.WithContext(ctx).Create(stockLevelDomain).Error
}
