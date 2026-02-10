package repository

import (
	"context"
	"putra4648/erp/internal/modules/stock_adjustment/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type stockAdjustmentRepository struct {
	db *gorm.DB
}

func NewStockAdjustmentRepository(db *gorm.DB) domain.StockAdjustmentRepository {
	return &stockAdjustmentRepository{db: db}
}

func (r *stockAdjustmentRepository) Save(ctx context.Context, adjustment *domain.StockAdjustment) error {
	return r.db.WithContext(ctx).Create(adjustment).Error
}

func (r *stockAdjustmentRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.StockAdjustment, error) {
	var adjustment domain.StockAdjustment
	err := r.db.WithContext(ctx).
		Preload("Warehouse").
		Preload("Items.Product").
		Preload("Items.Reason").
		First(&adjustment, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &adjustment, nil
}

func (r *stockAdjustmentRepository) FindAll(ctx context.Context, page, size int) ([]*domain.StockAdjustment, int64, error) {
	var adjustments []*domain.StockAdjustment
	var total int64

	model := r.db.WithContext(ctx).Model(&domain.StockAdjustment{})
	model.Count(&total)

	offset := (page - 1) * size
	err := model.
		Preload("Warehouse").
		Offset(offset).Limit(size).Find(&adjustments).Error
	if err != nil {
		return nil, 0, err
	}
	return adjustments, total, nil
}

func (r *stockAdjustmentRepository) Update(ctx context.Context, adjustment *domain.StockAdjustment) error {
	// Full update including items might be complex depending on use case.
	// For now, using Save which handles associations if configured.
	return r.db.WithContext(ctx).Save(adjustment).Error
}

func (r *stockAdjustmentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&domain.StockAdjustment{}, "id = ?", id).Error
}
