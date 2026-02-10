package repository

import (
	"context"
	productDomain "putra4648/erp/internal/modules/product/domain"
	"putra4648/erp/internal/modules/product/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(ctx context.Context, product *productDomain.Product) error
	FindByID(ctx context.Context, id uuid.UUID) (*productDomain.Product, error)
	FindAll(ctx context.Context, req *dto.ProductRequest) ([]*productDomain.Product, int64, error)
	Update(ctx context.Context, product *productDomain.Product) error
	Delete(ctx context.Context, id uuid.UUID) error
	FindBySKU(ctx context.Context, sku string) (*productDomain.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(ctx context.Context, product *productDomain.Product) error {
	return r.db.WithContext(ctx).Create(product).Error
}

func (r *productRepository) FindByID(ctx context.Context, id uuid.UUID) (*productDomain.Product, error) {
	var product productDomain.Product
	err := r.db.WithContext(ctx).Preload("Categories.Category").Preload("UOMs.UOM").Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) FindAll(ctx context.Context, req *dto.ProductRequest) ([]*productDomain.Product, int64, error) {
	var products []*productDomain.Product
	var total int64
	db := r.db.WithContext(ctx).Model(&productDomain.Product{}).Preload("Categories.Category").Preload("UOMs.UOM")

	if req.Name != "" {
		db = db.Where("name ILIKE ?", "%"+req.Name+"%")
	}

	db.Count(&total)

	if req.Page > 0 && req.Size > 0 {
		offset := (req.Page - 1) * req.Size
		db = db.Limit(req.Size).Offset(offset)
	}

	err := db.Find(&products).Error
	if err != nil {
		return nil, 0, err
	}
	return products, total, nil
}

func (r *productRepository) Update(ctx context.Context, product *productDomain.Product) error {
	tx := r.db.WithContext(ctx).Begin()
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

func (r *productRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&productDomain.Product{}, "id = ?", id).Error
}

func (r *productRepository) FindBySKU(ctx context.Context, sku string) (*productDomain.Product, error) {
	var product productDomain.Product
	err := r.db.WithContext(ctx).Where("sku = ?", sku).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}
