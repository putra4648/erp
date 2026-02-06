package service

import (
	"putra4648/erp/internal/modules/product/model"
	"putra4648/erp/internal/modules/product/repository"

	"github.com/google/uuid"
)

type ProductCommandService interface {
	CreateProduct(productDTO *model.ProductDTO) (*model.ProductResponse, error)
	UpdateProduct(id uuid.UUID, productDTO *model.ProductDTO) (*model.ProductResponse, error)
	DeleteProduct(id uuid.UUID) error
}

type productCommandService struct {
	productRepo repository.ProductRepository
}

func NewProductCommandService(productRepo repository.ProductRepository) ProductCommandService {
	return &productCommandService{productRepo: productRepo}
}

func (s *productCommandService) CreateProduct(productDTO *model.ProductDTO) (*model.ProductResponse, error) {
	// Check if SKU already exists
	existingProduct, err := s.productRepo.FindBySKU(productDTO.SKU)
	if err == nil && existingProduct != nil {
		return nil, &ProductError{Code: "DUPLICATE_SKU", Message: "SKU already exists"}
	}

	// Convert DTO to model
	product := productDTO.ToModel()

	// Create product in database
	err = s.productRepo.Create(product)
	if err != nil {
		return nil, &ProductError{Code: "DATABASE_ERROR", Message: "Failed to create product"}
	}

	// Return response
	return product.ToResponse(), nil
}

func (s *productCommandService) UpdateProduct(id uuid.UUID, productDTO *model.ProductDTO) (*model.ProductResponse, error) {
	// Find existing product
	existingProduct, err := s.productRepo.FindByID(id)
	if err != nil {
		return nil, &ProductError{Code: "NOT_FOUND", Message: "Product not found"}
	}

	// Check if SKU is being changed and already exists
	if productDTO.SKU != existingProduct.SKU {
		skuExists, err := s.productRepo.FindBySKU(productDTO.SKU)
		if err == nil && skuExists != nil {
			return nil, &ProductError{Code: "DUPLICATE_SKU", Message: "SKU already exists"}
		}
	}

	// Update product fields
	existingProduct.Name = productDTO.Name
	existingProduct.Description = productDTO.Description
	existingProduct.SKU = productDTO.SKU
	existingProduct.Price = productDTO.Price
	existingProduct.Cost = productDTO.Cost
	existingProduct.Quantity = productDTO.Quantity
	// Convert Category DTOs to models
	categories := make([]model.Category, len(productDTO.Category))
	for i, catDTO := range productDTO.Category {
		categories[i] = *catDTO.ToModel()
	}
	existingProduct.Category = categories
	// Convert UOM DTO to model
	existingProduct.UOM = *productDTO.UOM.ToModel()
	existingProduct.IsActive = productDTO.IsActive

	// Save updated product
	err = s.productRepo.Update(existingProduct)
	if err != nil {
		return nil, &ProductError{Code: "DATABASE_ERROR", Message: "Failed to update product"}
	}

	// Return response
	return existingProduct.ToResponse(), nil
}

func (s *productCommandService) DeleteProduct(id uuid.UUID) error {
	// Find product first to ensure it exists
	_, err := s.productRepo.FindByID(id)
	if err != nil {
		return &ProductError{Code: "NOT_FOUND", Message: "Product not found"}
	}

	// Delete product
	err = s.productRepo.Delete(id)
	if err != nil {
		return &ProductError{Code: "DATABASE_ERROR", Message: "Failed to delete product"}
	}

	return nil
}
