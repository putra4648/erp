package repository

import (
	"context"
	"putra4648/erp/internal/product/dto"

	"github.com/google/uuid"

	"putra4648/erp/internal/product/domain"
)

type ProductRepository interface {
	Create(ctx context.Context, product *domain.Product) error
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Product, error)
	FindAll(ctx context.Context, req *dto.ProductRequest) ([]*domain.Product, int64, error)
	Update(ctx context.Context, product *domain.Product) error
	Delete(ctx context.Context, id uuid.UUID) error
	FindBySKU(ctx context.Context, sku string) (*domain.Product, error)
}
