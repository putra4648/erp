package service

import (
	"context"
	"putra4648/erp/internal/modules/product/domain"
	"putra4648/erp/internal/modules/product/dto"
	"putra4648/erp/internal/modules/product/repository"
	sharedDto "putra4648/erp/internal/modules/shared/dto"

	"github.com/google/uuid"
)

type ProductQueryService interface {
	GetProductByID(ctx context.Context, id uuid.UUID) (*domain.ProductResponse, error)
	GetAllProducts(ctx context.Context, req *dto.ProductRequest) (*sharedDto.PaginationResponse[*domain.ProductResponse], error)
}

type productQueryService struct {
	productRepo repository.ProductRepository
}

func NewProductQueryService(productRepo repository.ProductRepository) ProductQueryService {
	return &productQueryService{productRepo: productRepo}
}

func (s *productQueryService) GetProductByID(ctx context.Context, id uuid.UUID) (*domain.ProductResponse, error) {
	// Find product in database
	product, err := s.productRepo.FindByID(ctx, id)
	if err != nil {
		return nil, &ProductError{Code: "NOT_FOUND", Message: "Product not found"}
	}

	// Return response
	return product.ToResponse(), nil
}

func (s *productQueryService) GetAllProducts(ctx context.Context, req *dto.ProductRequest) (*sharedDto.PaginationResponse[*domain.ProductResponse], error) {
	// Find all products in database
	products, total, err := s.productRepo.FindAll(ctx, req)
	if err != nil {
		return nil, &ProductError{Code: "DATABASE_ERROR", Message: "Failed to retrieve products"}
	}

	// Convert to response format
	responses := make([]*domain.ProductResponse, len(products))
	for i, product := range products {
		responses[i] = product.ToResponse()
	}

	return &sharedDto.PaginationResponse[*domain.ProductResponse]{
		Items: responses,
		Total: total,
		Page:  req.Page,
		Size:  req.Size,
	}, nil
}
