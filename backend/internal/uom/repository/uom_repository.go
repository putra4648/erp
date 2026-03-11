package repository

import (
	"context"
	sharedDto "putra4648/erp/internal/shared/dto"
	"putra4648/erp/internal/uom/domain"
	"putra4648/erp/internal/uom/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type uomRepository struct {
	db *gorm.DB
}

func NewUOMRepository(db *gorm.DB) UOMRepository {
	return &uomRepository{db: db}
}

func (r *uomRepository) Create(ctx context.Context, uom *domain.UOM) error {
	return r.db.WithContext(ctx).Create(uom).Error
}

func (r *uomRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.UOM, error) {
	var uom domain.UOM
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&uom).Error
	if err != nil {
		return nil, err
	}
	return &uom, nil
}

func (r *uomRepository) FindAll(ctx context.Context, pagination *sharedDto.PaginationRequest, req *dto.UOMDTO) ([]*domain.UOM, int64, error) {
	var uoms []*domain.UOM
	var total int64
	db := r.db.WithContext(ctx).Model(&domain.UOM{})

	if req.Name != "" {
		db = db.Where("name ILIKE ?", "%"+req.Name+"%")
	}

	db.Count(&total)

	if pagination.Page > 0 && pagination.Size > 0 {
		offset := (pagination.Page - 1) * pagination.Size
		db = db.Limit(pagination.Size).Offset(offset)
	}

	err := db.Find(&uoms).Error
	if err != nil {
		return nil, 0, err
	}
	return uoms, total, nil
}

func (r *uomRepository) Update(ctx context.Context, uom *domain.UOM) error {
	return r.db.WithContext(ctx).Save(uom).Error
}

func (r *uomRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&domain.UOM{}, "id = ?", id).Error
}
