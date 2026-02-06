package repository

import (
	"putra4648/erp/internal/modules/product/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UOMRepository interface {
	Create(uom *model.UOM) error
	FindByID(id uuid.UUID) (*model.UOM, error)
	FindAll() ([]*model.UOM, error)
	Update(uom *model.UOM) error
	Delete(id uuid.UUID) error
}

type uomRepository struct {
	db *gorm.DB
}

func NewUOMRepository(db *gorm.DB) UOMRepository { // Corrected return type
	return &uomRepository{db: db} // Corrected struct name
}

func (r *uomRepository) Create(uom *model.UOM) error {
	return r.db.Create(uom).Error
}

func (r *uomRepository) FindByID(id uuid.UUID) (*model.UOM, error) {
	var uom model.UOM
	err := r.db.Where("id = ?", id).First(&uom).Error
	if err != nil {
		return nil, err
	}
	return &uom, nil
}

func (r *uomRepository) FindAll() ([]*model.UOM, error) {
	var uoms []*model.UOM
	err := r.db.Find(&uoms).Error
	if err != nil {
		return nil, err
	}
	return uoms, nil
}

func (r *uomRepository) Update(uom *model.UOM) error {
	return r.db.Save(uom).Error
}

func (r *uomRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&model.UOM{}, "id = ?", id).Error // Corrected model to delete
}
