package service

import (
	"context"
	categoryRepository "putra4648/erp/internal/category/repository"
	productDto "putra4648/erp/internal/product/dto"
	"putra4648/erp/internal/product/mapper"
	"putra4648/erp/internal/product/repository"
	productRepository "putra4648/erp/internal/product/repository"
	supplierRepository "putra4648/erp/internal/supplier/repository"
	uomRepository "putra4648/erp/internal/uom/repository"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type productCommandService struct {
	productRepo  repository.ProductRepository
	categoryRepo categoryRepository.CategoryRepository
	uomRepo      uomRepository.UOMRepository
	supplierRepo supplierRepository.SupplierRepository
}

func NewProductCommandService(
	productRepo productRepository.ProductRepository,
	categoryRepo categoryRepository.CategoryRepository,
	uomRepo uomRepository.UOMRepository,
	supplierRepo supplierRepository.SupplierRepository,
	logger *zap.Logger,
) ProductCommandService {
	return &productCommandService{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
		uomRepo:      uomRepo,
		supplierRepo: supplierRepo,
	}
}

func (s *productCommandService) CreateProduct(ctx context.Context, productDTO *productDto.ProductDTO) (*productDto.ProductDTO, error) {
	// Check if SKU already exists
	if _, err := s.productRepo.FindBySKU(ctx, productDTO.SKU); err == nil {
		return nil, &ProductError{Code: "DUPLICATE_SKU", Message: "SKU already exists"}
	}

	// Validate UOMs
	for _, uomDTO := range productDTO.UOMs {
		uomID, err := uuid.Parse(uomDTO.ID)
		if err != nil {
			return nil, &ProductError{Code: "INVALID_ID", Message: "Invalid UOM ID"}
		}
		if _, err := s.uomRepo.FindByID(ctx, uomID); err != nil {
			return nil, &ProductError{Code: "NOT_FOUND", Message: "UOM not found"}
		}
	}

	// Validate Categories
	for _, catDTO := range productDTO.Categories {
		catID, err := uuid.Parse(catDTO.ID)
		if err != nil {
			return nil, &ProductError{Code: "INVALID_ID", Message: "Invalid Category ID"}
		}
		if _, err := s.categoryRepo.FindByID(ctx, catID); err != nil {
			return nil, &ProductError{Code: "NOT_FOUND", Message: "Category not found"}
		}
	}

	// Validate Supplier
	supplierID, err := uuid.Parse(productDTO.SupplierID)
	if err != nil {
		return nil, &ProductError{Code: "INVALID_ID", Message: "Invalid Supplier ID"}
	}
	if _, err := s.supplierRepo.FindByID(ctx, supplierID); err != nil {
		return nil, &ProductError{Code: "NOT_FOUND", Message: "Supplier not found"}
	}

	product := mapper.ToProduct(productDTO)

	// Create product in database
	if err := s.productRepo.Create(ctx, product); err != nil {
		return nil, &ProductError{Code: "DATABASE_ERROR", Message: "Failed to create product"}
	}

	// Reload product to get all fields populated, including associations
	createdProduct, err := s.productRepo.FindByID(ctx, product.ID)
	if err != nil {
		return nil, &ProductError{Code: "DATABASE_ERROR", Message: "Failed to retrieve created product"}
	}

	return mapper.ToProductDTO(createdProduct), nil
}

func (s *productCommandService) UpdateProduct(ctx context.Context, id uuid.UUID, productDTO *productDto.ProductDTO) (*productDto.ProductDTO, error) {
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
	for _, uomDTO := range productDTO.UOMs {
		uomID, err := uuid.Parse(uomDTO.ID)
		if err != nil {
			return nil, &ProductError{Code: "INVALID_ID", Message: "Invalid UOM ID"}
		}
		if _, err := s.uomRepo.FindByID(ctx, uomID); err != nil {
			return nil, &ProductError{Code: "NOT_FOUND", Message: "UOM not found"}
		}
	}

	// Validate Categories
	for _, catDTO := range productDTO.Categories {
		catID, err := uuid.Parse(catDTO.ID)
		if err != nil {
			return nil, &ProductError{Code: "INVALID_ID", Message: "Invalid Category ID"}
		}
		if _, err := s.categoryRepo.FindByID(ctx, catID); err != nil {
			return nil, &ProductError{Code: "NOT_FOUND", Message: "Category not found"}
		}
	}

	// Validate Supplier
	supplierID, err := uuid.Parse(productDTO.SupplierID)
	if err != nil {
		return nil, &ProductError{Code: "INVALID_ID", Message: "Invalid Supplier ID"}
	}
	if _, err := s.supplierRepo.FindByID(ctx, supplierID); err != nil {
		return nil, &ProductError{Code: "NOT_FOUND", Message: "Supplier not found"}
	}

	updatedProduct := mapper.ToProduct(productDTO)
	updatedProduct.ID = id

	// Save updated product
	if err := s.productRepo.Update(ctx, updatedProduct); err != nil {
		return nil, &ProductError{Code: "DATABASE_ERROR", Message: "Failed to update product"}
	}

	// Reload product to get all fields populated, including associations
	reloadedProduct, err := s.productRepo.FindByID(ctx, id)
	if err != nil {
		return nil, &ProductError{Code: "DATABASE_ERROR", Message: "Failed to retrieve updated product"}
	}

	return mapper.ToProductDTO(reloadedProduct), nil
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
