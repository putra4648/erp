package repository

import (
	"context"
	"putra4648/erp/internal/category/domain"
	"putra4648/erp/internal/category/dto"
	sharedDto "putra4648/erp/internal/shared/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(ctx context.Context, category *domain.Category) error {
	return r.db.WithContext(ctx).Create(category).Error
}

func (r *categoryRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Category, error) {
	var category domain.Category
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) FindAll(ctx context.Context, pagination *sharedDto.PaginationRequest, req *dto.CategoryDTO) ([]*domain.Category, int64, error) {
	var categories []*domain.Category
	var total int64

	db := r.db.WithContext(ctx).Model(&domain.Category{})

	if req.Name != "" {
		db = db.Where("name ILIKE ?", "%"+req.Name+"%")
	}

	db.Count(&total)

	if pagination.Page > 0 && pagination.Size > 0 {
		offset := (pagination.Page - 1) * pagination.Size
		db = db.Offset(offset).Limit(pagination.Size)
	}

	err := db.Find(&categories).Error
	if err != nil {
		return nil, 0, err
	}
	return categories, total, nil
}

func (r *categoryRepository) Update(ctx context.Context, category *domain.Category) error {
	return r.db.WithContext(ctx).Save(category).Error
}

func (r *categoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&domain.Category{}, "id = ?", id).Error
}
