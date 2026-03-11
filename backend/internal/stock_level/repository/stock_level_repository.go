package repository

import (
	"context"
	sharedDto "putra4648/erp/internal/shared/dto"
	"putra4648/erp/internal/stock_level/domain"
	"putra4648/erp/internal/stock_level/dto"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type stockLevelRepository struct {
	db *gorm.DB
}

func NewStockLevelRepository(db *gorm.DB) StockLevelRepository {
	return &stockLevelRepository{db: db}
}

func (r *stockLevelRepository) FindByProductAndWarehouse(ctx context.Context, dto *dto.StockLevelDto) (*domain.StockLevel, error) {
	var stockLevel domain.StockLevel
	err := r.db.WithContext(ctx).
		Where("product_id = ? AND warehouse_id = ?", dto.ProductID, dto.WarehouseID).
		First(&stockLevel).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &stockLevel, nil
}

func (r *stockLevelRepository) FindByProductAndWarehouseWithPreload(ctx context.Context, dto *dto.StockLevelDto) (*domain.StockLevel, error) {
	var stockLevel domain.StockLevel
	err := r.db.WithContext(ctx).
		Preload("Product").Preload("Warehouse").
		Where("product_id = ? AND warehouse_id = ?", dto.ProductID, dto.WarehouseID).
		First(&stockLevel).Error
	if err != nil {
		return nil, err
	}
	return &stockLevel, nil
}

func (r *stockLevelRepository) UpdateQuantity(ctx context.Context, dto *dto.StockLevelDto) error {
	var stockLevel domain.StockLevel
	err := r.db.WithContext(ctx).
		Where("product_id = ? AND warehouse_id = ?", dto.ProductID, dto.WarehouseID).
		First(&stockLevel).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create new stock level
			stockLevel = domain.StockLevel{
				ID:          uuid.New(),
				ProductID:   uuid.MustParse(*dto.ProductID),
				WarehouseID: uuid.MustParse(*dto.WarehouseID),
				Quantity:    dto.Quantity,
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
	stockLevel.Quantity = dto.Quantity
	stockLevel.LastUpdated = time.Now()
	return r.db.WithContext(ctx).Save(&stockLevel).Error
}

func (r *stockLevelRepository) FindStockLevels(ctx context.Context, pagination *sharedDto.PaginationRequest, dto *dto.StockLevelDto) ([]*domain.StockLevel, int64, error) {
	var stockLevels []*domain.StockLevel
	var total int64

	db := r.db.WithContext(ctx).Model(&domain.StockLevel{}).Preload("Product").Preload("Warehouse")

	if dto.WarehouseID != nil {
		db = db.Where("warehouse_id = ?", dto.WarehouseID)
	}
	if dto.ProductID != nil {
		db = db.Where("product_id = ?", dto.ProductID)
	}

	if dto.Name != "" {
		db = db.Joins("JOIN products ON products.id = stock_levels.product_id").
			Where("products.name ILIKE ?", "%"+dto.Name+"%")
	}

	db.Count(&total)

	if pagination.Page > 0 && pagination.Size > 0 {
		offset := (pagination.Page - 1) * pagination.Size
		db = db.Offset(offset).Limit(pagination.Size)
	}

	err := db.Find(&stockLevels).Error
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

func (r *stockLevelRepository) Create(ctx context.Context, stockLevel *domain.StockLevel) error {
	if stockLevel.ID == uuid.Nil {
		stockLevel.ID = uuid.New()
	}
	stockLevel.LastUpdated = time.Now()
	return r.db.WithContext(ctx).Create(stockLevel).Error
}
