package service

import (
	"context"
	"putra4648/erp/internal/product/dto"
	productDto "putra4648/erp/internal/product/dto"
	sharedDto "putra4648/erp/internal/shared/dto"

	"github.com/google/uuid"
)

type ProductCommandService interface {
	CreateProduct(ctx context.Context, req *productDto.ProductDTO) (*productDto.ProductDTO, error)
	UpdateProduct(ctx context.Context, id uuid.UUID, productDTO *productDto.ProductDTO) (*productDto.ProductDTO, error)
	DeleteProduct(ctx context.Context, id uuid.UUID) error
}

type ProductQueryService interface {
	GetProductByID(ctx context.Context, id uuid.UUID) (*dto.ProductDTO, error)
	GetAllProducts(ctx context.Context, req *dto.ProductRequest) (*sharedDto.PaginationResponse[*dto.ProductDTO], error)
}
