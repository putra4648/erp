package repository

import (
	"context"
	"putra4648/erp/internal/modules/inventory/supplier/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type supplierRepository struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) domain.SupplierRepository {
	return &supplierRepository{db: db}
}

func (r *supplierRepository) Save(ctx context.Context, supplier *domain.Supplier) error {
	return r.db.WithContext(ctx).Create(supplier).Error
}

func (r *supplierRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Supplier, error) {
	var supplier domain.Supplier
	err := r.db.WithContext(ctx).First(&supplier, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &supplier, nil
}

func (r *supplierRepository) FindAll(ctx context.Context, page, size int) ([]*domain.Supplier, int64, error) {
	var suppliers []*domain.Supplier
	var total int64

	r.db.WithContext(ctx).Model(&domain.Supplier{}).Count(&total)

	offset := (page - 1) * size
	err := r.db.WithContext(ctx).Offset(offset).Limit(size).Find(&suppliers).Error
	if err != nil {
		return nil, 0, err
	}
	return suppliers, total, nil
}

func (r *supplierRepository) Update(ctx context.Context, supplier *domain.Supplier) error {
	return r.db.WithContext(ctx).Save(supplier).Error
}

func (r *supplierRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&domain.Supplier{}, "id = ?", id).Error
}
