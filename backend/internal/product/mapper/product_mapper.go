package mapper

import (
	categoryMapper "putra4648/erp/internal/category/mapper"
	"putra4648/erp/internal/product/domain"
	"putra4648/erp/internal/product/dto"
	uomMapper "putra4648/erp/internal/uom/mapper"

	"github.com/google/uuid"
)

func ToProduct(productDTO *dto.ProductDTO) *domain.Product {
	return &domain.Product{
		Name: productDTO.Name,
		// Description: productDTO.Description,
		Price:      productDTO.Price,
		SKU:        productDTO.SKU,
		SupplierID: uuid.MustParse(productDTO.SupplierID),
		Categories: categoryMapper.ToCategories(productDTO.Categories),
		UOMs:       uomMapper.ToUOMs(productDTO.UOMs),
	}
}

func ToProductDTO(product *domain.Product) *dto.ProductDTO {
	return &dto.ProductDTO{
		ID:   product.ID.String(),
		Name: product.Name,
		// Description: product.Description,
		Price:      product.Price,
		SKU:        product.SKU,
		SupplierID: product.SupplierID.String(),
		Categories: categoryMapper.ToCategoryDTOs(product.Categories),
		UOMs:       uomMapper.ToUOMDTOs(product.UOMs),
	}
}

func ToProductDTOs(products []*domain.Product) []*dto.ProductDTO {
	productDTOs := make([]*dto.ProductDTO, len(products))
	for i, product := range products {
		productDTOs[i] = ToProductDTO(product)
	}
	return productDTOs
}
