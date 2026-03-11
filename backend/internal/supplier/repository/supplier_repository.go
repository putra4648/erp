package repository

import (
	"context"
	sharedDto "putra4648/erp/internal/shared/dto"
	"putra4648/erp/internal/supplier/domain"
	"putra4648/erp/internal/supplier/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type supplierRepository struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) SupplierRepository {
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

func (r *supplierRepository) FindAll(ctx context.Context, pagination *sharedDto.PaginationRequest, req *dto.SupplierDTO) ([]*domain.Supplier, int64, error) {
	var suppliers []*domain.Supplier
	var total int64

	db := r.db.WithContext(ctx).Model(&domain.Supplier{})

	if req.Name != "" {
		db = db.Where("name ILIKE ?", "%"+req.Name+"%")
	}

	db.Count(&total)

	if pagination.Page > 0 && pagination.Size > 0 {
		offset := (pagination.Page - 1) * pagination.Size
		db = db.Offset(offset).Limit(pagination.Size)
	}

	err := db.Find(&suppliers).Error
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
