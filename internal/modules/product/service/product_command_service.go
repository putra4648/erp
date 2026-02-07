package service

import (
	"context"
	categoryDomain "putra4648/erp/internal/modules/category/domain"
	categoryRepository "putra4648/erp/internal/modules/category/repository"
	productDomain "putra4648/erp/internal/modules/product/domain"
	productRepository "putra4648/erp/internal/modules/product/repository"
	uomDomain "putra4648/erp/internal/modules/uom/domain" // Added import
	uomRepository "putra4648/erp/internal/modules/uom/repository"

	"github.com/google/uuid"
)

type ProductCommandService interface {
	CreateProduct(ctx context.Context, req *productDomain.ProductDTO) (*productDomain.ProductResponse, error)
	UpdateProduct(ctx context.Context, id uuid.UUID, productDTO *productDomain.ProductDTO) (*productDomain.ProductResponse, error)
	DeleteProduct(ctx context.Context, id uuid.UUID) error
}

type productCommandService struct {
	productRepo  productRepository.ProductRepository
	categoryRepo categoryRepository.CategoryRepository
	uomRepo      uomRepository.UOMRepository
}

func NewProductCommandService(
	productRepo productRepository.ProductRepository,
	categoryRepo categoryRepository.CategoryRepository,
	uomRepo uomRepository.UOMRepository,
) ProductCommandService {
	return &productCommandService{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
		uomRepo:      uomRepo,
	}
}

func (s *productCommandService) CreateProduct(ctx context.Context, productDTO *productDomain.ProductDTO) (*productDomain.ProductResponse, error) {
	// Check if SKU already exists
	if _, err := s.productRepo.FindBySKU(ctx, productDTO.SKU); err == nil {
		return nil, &ProductError{Code: "DUPLICATE_SKU", Message: "SKU already exists"}
	}

	// Validate UOMs
	uoms := make([]*uomDomain.UOM, len(productDTO.UOMIDs))
	for i, uomID := range productDTO.UOMIDs {
		uom, err := s.uomRepo.FindByID(ctx, uomID)
		if err != nil {
			return nil, &ProductError{Code: "NOT_FOUND", Message: "UOM not found"}
		}
		uoms[i] = uom
	}

	// Validate Categories
	categories := make([]*categoryDomain.Category, len(productDTO.CategoryIDs))
	for i, catID := range productDTO.CategoryIDs {
		category, err := s.categoryRepo.FindByID(ctx, catID)
		if err != nil {
			return nil, &ProductError{Code: "NOT_FOUND", Message: "Category not found"}
		}
		categories[i] = category
	}

	// Convert DTO to model
	product := productDTO.ToModel()
	product.UOMs = uoms
	product.Categories = categories

	// Create product in database
	if err := s.productRepo.Create(ctx, product); err != nil {
		return nil, &ProductError{Code: "DATABASE_ERROR", Message: "Failed to create product"}
	}

	// Reload product to get all fields populated, including associations
	createdProduct, err := s.productRepo.FindByID(ctx, product.ID)
	if err != nil {
		return nil, &ProductError{Code: "DATABASE_ERROR", Message: "Failed to retrieve created product"}
	}

	// Return response
	return createdProduct.ToResponse(), nil
}

func (s *productCommandService) UpdateProduct(ctx context.Context, id uuid.UUID, productDTO *productDomain.ProductDTO) (*productDomain.ProductResponse, error) {
	// Find existing product
	existingProduct, err := s.productRepo.FindByID(ctx, id)
	if err != nil {
		return nil, &ProductError{Code: "NOT_FOUND", Message: "Product not found"}
	}

	// Check if SKU is being changed and already exists
	if productDTO.SKU != existingProduct.SKU {
		if _, err := s.productRepo.FindBySKU(ctx, productDTO.SKU); err == nil {
			return nil, &ProductError{Code: "DUPLICATE_SKU", Message: "SKU already exists"}
		}
	}

	// Validate UOMs
	uoms := make([]*uomDomain.UOM, len(productDTO.UOMIDs))
	for i, uomID := range productDTO.UOMIDs {
		uom, err := s.uomRepo.FindByID(ctx, uomID)
		if err != nil {
			return nil, &ProductError{Code: "NOT_FOUND", Message: "UOM not found"}
		}
		uoms[i] = uom
	}

	// Validate Categories
	categories := make([]*categoryDomain.Category, len(productDTO.CategoryIDs))
	for i, catID := range productDTO.CategoryIDs {
		category, err := s.categoryRepo.FindByID(ctx, catID)
		if err != nil {
			return nil, &ProductError{Code: "NOT_FOUND", Message: "Category not found"}
		}
		categories[i] = category
	}

	// Update product fields
	existingProduct.Name = productDTO.Name
	existingProduct.Description = productDTO.Description
	existingProduct.SKU = productDTO.SKU
	existingProduct.Price = productDTO.Price
	existingProduct.Cost = productDTO.Cost
	existingProduct.Quantity = productDTO.Quantity
	existingProduct.IsActive = productDTO.IsActive
	existingProduct.Categories = categories
	existingProduct.UOMs = uoms

	// Save updated product
	if err := s.productRepo.Update(ctx, existingProduct); err != nil {
		return nil, &ProductError{Code: "DATABASE_ERROR", Message: "Failed to update product"}
	}

	// Reload product to get all fields populated, including associations
	updatedProduct, err := s.productRepo.FindByID(ctx, id)
	if err != nil {
		return nil, &ProductError{Code: "DATABASE_ERROR", Message: "Failed to retrieve updated product"}
	}

	// Return response
	return updatedProduct.ToResponse(), nil
}

func (s *productCommandService) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	// Find product first to ensure it exists
	_, err := s.productRepo.FindByID(ctx, id)
	if err != nil {
		return &ProductError{Code: "NOT_FOUND", Message: "Product not found"}
	}

	// Delete product
	err = s.productRepo.Delete(ctx, id)
	if err != nil {
		return &ProductError{Code: "DATABASE_ERROR", Message: "Failed to delete product"}
	}

	return nil
}
