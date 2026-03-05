package service

import (
	"context"
	"putra4648/erp/internal/product/dto"
	"putra4648/erp/internal/product/mapper"
	"putra4648/erp/internal/product/repository"
	sharedDto "putra4648/erp/internal/shared/dto"

	"github.com/google/uuid"
)

type productQueryService struct {
	productRepo repository.ProductRepository
}

func NewProductQueryService(productRepo repository.ProductRepository) ProductQueryService {
	return &productQueryService{productRepo: productRepo}
}

func (s *productQueryService) GetProductByID(ctx context.Context, id uuid.UUID) (*dto.ProductDTO, error) {
	// Find product in database
	product, err := s.productRepo.FindByID(ctx, id)
	if err != nil {
		return nil, &ProductError{Code: "NOT_FOUND", Message: "Product not found"}
	}

	return mapper.ToProductDTO(product), nil
}

func (s *productQueryService) GetAllProducts(ctx context.Context, req *dto.ProductRequest) (*sharedDto.PaginationResponse[*dto.ProductDTO], error) {
	// Find all products in database
	products, total, err := s.productRepo.FindAll(ctx, req)
	if err != nil {
		return nil, &ProductError{Code: "DATABASE_ERROR", Message: "Failed to retrieve products"}
	}

	// Convert to response format
	responses := mapper.ToProductDTOs(products)

	return &sharedDto.PaginationResponse[*dto.ProductDTO]{
		Items: responses,
		Total: total,
		Page:  req.Page,
		Size:  req.Size,
	}, nil
}
