package repository

import (
	"context"
	sharedDto "putra4648/erp/internal/shared/dto"
	"putra4648/erp/internal/stock_adjustment/domain"
	"putra4648/erp/internal/stock_adjustment/dto"

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

func (r *adjustmentReasonRepository) FindAll(ctx context.Context, pagination *sharedDto.PaginationRequest, dto *dto.AdjustmentReasonDto) ([]*domain.AdjustmentReason, int64, error) {
	var reasons []*domain.AdjustmentReason
	var total int64
	db := r.db.WithContext(ctx).Model(&domain.AdjustmentReason{})

	if pagination.Page > 0 && pagination.Size > 0 {
		offset := (pagination.Page - 1) * pagination.Size
		db = db.Offset(offset).Limit(pagination.Size)
	}

	err := db.Find(&reasons, "code LIKE ?", "%"+dto.Code+"%").Error
	if err != nil {
		return nil, 0, err
	}

	return reasons, total, nil
}

func (r *adjustmentReasonRepository) Update(ctx context.Context, reason *domain.AdjustmentReason) error {
	return r.db.WithContext(ctx).Save(reason).Error
}

func (r *adjustmentReasonRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&domain.AdjustmentReason{}, "id = ?", id).Error
}
