package repository

import (
	productModel "putra4648/erp/internal/modules/product/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *productModel.Product) error
	FindByID(id uuid.UUID) (*productModel.Product, error)
	FindAll() ([]*productModel.Product, error)
	Update(product *productModel.Product) error
	Delete(id uuid.UUID) error
	FindBySKU(sku string) (*productModel.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product *productModel.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) FindByID(id uuid.UUID) (*productModel.Product, error) {
	var product productModel.Product
	err := r.db.Preload("Categories").Preload("UOMs").Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) FindAll() ([]*productModel.Product, error) {
	var products []*productModel.Product
	err := r.db.Preload("Categories").Preload("UOMs").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepository) Update(product *productModel.Product) error {
	tx := r.db.Begin()
	if err := tx.Model(&product).Association("Categories").Replace(product.Categories); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&product).Association("UOMs").Replace(product.UOMs); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Save(product).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *productRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&productModel.Product{}, "id = ?", id).Error
}

func (r *productRepository) FindBySKU(sku string) (*productModel.Product, error) {
	var product productModel.Product
	err := r.db.Where("sku = ?", sku).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}