package service

import (
	"putra4648/erp/internal/modules/product/model"
	"putra4648/erp/internal/modules/product/repository"

	"github.com/google/uuid"
)

type ProductQueryService interface {
	GetProductByID(id uuid.UUID) (*model.ProductResponse, error)
	GetAllProducts() ([]*model.ProductResponse, error)
}

type productQueryService struct {
	productRepo repository.ProductRepository
}

func NewProductQueryService(productRepo repository.ProductRepository) ProductQueryService {
	return &productQueryService{productRepo: productRepo}
}

func (s *productQueryService) GetProductByID(id uuid.UUID) (*model.ProductResponse, error) {
	// Find product in database
	product, err := s.productRepo.FindByID(id)
	if err != nil {
		return nil, &ProductError{Code: "NOT_FOUND", Message: "Product not found"}
	}

	// Return response
	return product.ToResponse(), nil
}

func (s *productQueryService) GetAllProducts() ([]*model.ProductResponse, error) {
	// Find all products in database
	products, err := s.productRepo.FindAll()
	if err != nil {
		return nil, &ProductError{Code: "DATABASE_ERROR", Message: "Failed to retrieve products"}
	}

	// Convert to response format
	responses := make([]*model.ProductResponse, len(products))
	for i, product := range products {
		responses[i] = product.ToResponse()
	}

	return responses, nil
}
