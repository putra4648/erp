package repository

import (
	"context"
	"putra4648/erp/internal/modules/product/domain"
	"putra4648/erp/internal/modules/product/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(ctx context.Context, product *domain.Product) error {
	return r.db.WithContext(ctx).Create(product).Error
}

func (r *productRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Product, error) {
	var product domain.Product
	err := r.db.WithContext(ctx).Preload("Categories").Preload("UOMs").Preload("Supplier").Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) FindAll(ctx context.Context, req *dto.ProductRequest) ([]*domain.Product, int64, error) {
	var products []*domain.Product
	var total int64
	db := r.db.WithContext(ctx).Model(&domain.Product{}).Preload("Categories").Preload("UOMs").Preload("Supplier")

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

func (r *productRepository) Update(ctx context.Context, product *domain.Product) error {
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
	return r.db.WithContext(ctx).Delete(&domain.Product{}, "id = ?", id).Error
}

func (r *productRepository) FindBySKU(ctx context.Context, sku string) (*domain.Product, error) {
	var product domain.Product
	err := r.db.WithContext(ctx).Where("sku = ?", sku).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}
