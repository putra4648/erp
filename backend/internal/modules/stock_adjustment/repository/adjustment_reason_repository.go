package repository

import (
	"context"
	"putra4648/erp/internal/modules/stock_adjustment/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type adjustmentReasonRepository struct {
	db *gorm.DB
}

func NewAdjustmentReasonRepository(db *gorm.DB) AdjustmentReasonRepository {
	return &adjustmentReasonRepository{db: db}
}

func (r *adjustmentReasonRepository) Save(ctx context.Context, reason *domain.AdjustmentReason) error {
	return r.db.WithContext(ctx).Create(reason).Error
}

func (r *adjustmentReasonRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.AdjustmentReason, error) {
	var reason domain.AdjustmentReason
	err := r.db.WithContext(ctx).First(&reason, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &reason, nil
}

func (r *adjustmentReasonRepository) FindAll(ctx context.Context) ([]*domain.AdjustmentReason, error) {
	var reasons []*domain.AdjustmentReason
	err := r.db.WithContext(ctx).Find(&reasons).Error
	if err != nil {
		return nil, err
	}
	return reasons, nil
}

func (r *adjustmentReasonRepository) Update(ctx context.Context, reason *domain.AdjustmentReason) error {
	return r.db.WithContext(ctx).Save(reason).Error
}

func (r *adjustmentReasonRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&domain.AdjustmentReason{}, "id = ?", id).Error
}
