package repository

import (
	"context"
	"putra4648/erp/internal/product/domain"
	"putra4648/erp/internal/product/dto"
	sharedDto "putra4648/erp/internal/shared/dto"

	"github.com/google/uuid"
)

type ProductRepository interface {
	Create(ctx context.Context, product *domain.Product) error
	Update(ctx context.Context, product *domain.Product) error
	Delete(ctx context.Context, id uuid.UUID) error

	FindByID(ctx context.Context, id uuid.UUID) (*domain.Product, error)
	FindAll(ctx context.Context, pagination *sharedDto.PaginationRequest, req *dto.ProductDTO) ([]*domain.Product, int64, error)
	FindBySKU(ctx context.Context, sku string) (*domain.Product, error)
}
