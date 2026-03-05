package repository

import (
	"context"
	"putra4648/erp/internal/stock_adjustment/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type stockAdjustmentRepository struct {
	db *gorm.DB
}

func NewStockAdjustmentRepository(db *gorm.DB) StockAdjustmentRepository {
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
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Use OMIT to not update created_by or other protected fields if necessary
		if err := tx.Save(adjustment).Error; err != nil {
			return err
		}
		// Replace items association to handle deletions of items not in the list
		return tx.Model(adjustment).Association("Items").Replace(adjustment.Items)
	})
}

func (r *stockAdjustmentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&domain.StockAdjustment{}, "id = ?", id).Error
}
