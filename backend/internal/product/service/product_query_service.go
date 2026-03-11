package service

import (
	"context"
	"putra4648/erp/internal/product/dto"
	"putra4648/erp/internal/product/mapper"
	"putra4648/erp/internal/product/repository"
	sharedDto "putra4648/erp/internal/shared/dto"
	"putra4648/erp/internal/shared/errors"

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
		return nil, &errors.ErrorDto{Code: "NOT_FOUND", Message: "Product not found"}
	}

	return mapper.ToProductDTO(product), nil
}

func (s *productQueryService) GetAllProducts(ctx context.Context, pagination *sharedDto.PaginationRequest, req *dto.ProductDTO) (*sharedDto.PaginationResponse[*dto.ProductDTO], error) {
	// Find all products in database
	products, total, err := s.productRepo.FindAll(ctx, pagination, req)
	if err != nil {
		return nil, &errors.ErrorDto{Code: "DATABASE_ERROR", Message: "Failed to retrieve products"}
	}

	// Convert to response format
	responses := mapper.ToProductDTOs(products)

	return &sharedDto.PaginationResponse[*dto.ProductDTO]{
		Items: responses,
		Total: total,
		Page:  pagination.Page,
		Size:  pagination.Size,
	}, nil
}
