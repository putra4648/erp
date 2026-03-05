package repository

import (
	"context"
	"putra4648/erp/internal/warehouse/domain"
	"putra4648/erp/internal/warehouse/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WarehouseRepository interface {
	Save(ctx context.Context, warehouse *domain.Warehouse) error
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Warehouse, error)
	FindAll(ctx context.Context, req *dto.WarehouseFindAllRequest) ([]*domain.Warehouse, int64, error)
	Update(ctx context.Context, warehouse *domain.Warehouse) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type warehouseRepository struct {
	db *gorm.DB
}

func NewWarehouseRepository(db *gorm.DB) WarehouseRepository {
	return &warehouseRepository{db: db}
}

func (r *warehouseRepository) Save(ctx context.Context, warehouse *domain.Warehouse) error {
	return r.db.WithContext(ctx).Create(warehouse).Error
}

func (r *warehouseRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Warehouse, error) {
	var warehouse domain.Warehouse
	err := r.db.WithContext(ctx).First(&warehouse, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &warehouse, nil
}

func (r *warehouseRepository) FindAll(ctx context.Context, req *dto.WarehouseFindAllRequest) ([]*domain.Warehouse, int64, error) {
	var warehouses []*domain.Warehouse
	var total int64

	db := r.db.WithContext(ctx).Model(&domain.Warehouse{})

	if req.Name != "" {
		db = db.Where("name ILIKE ?", "%"+req.Name+"%")
	}

	db.Count(&total)

	if req.Page > 0 && req.Size > 0 {
		offset := (req.Page - 1) * req.Size
		db = db.Limit(req.Size).Offset(offset)
	}

	err := db.Find(&warehouses).Error
	if err != nil {
		return nil, 0, err
	}
	return warehouses, total, nil
}

func (r *warehouseRepository) Update(ctx context.Context, warehouse *domain.Warehouse) error {
	return r.db.WithContext(ctx).Save(warehouse).Error
}

func (r *warehouseRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&domain.Warehouse{}, "id = ?", id).Error
}
