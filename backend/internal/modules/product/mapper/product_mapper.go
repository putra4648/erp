package mapper

import (
	categoryMapper "putra4648/erp/internal/modules/category/mapper"
	"putra4648/erp/internal/modules/product/domain"
	"putra4648/erp/internal/modules/product/dto"
	uomMapper "putra4648/erp/internal/modules/uom/mapper"
)

func ToProduct(productDTO *dto.ProductDTO) *domain.Product {
	return &domain.Product{
		Name:        productDTO.Name,
		Description: productDTO.Description,
		Price:       productDTO.Price,
		SKU:         productDTO.SKU,
		Categories:  categoryMapper.ToCategories(productDTO.Categories),
		UOMs:        uomMapper.ToUOMs(productDTO.UOMs),
	}
}

func ToProductDTO(product *domain.Product) *dto.ProductDTO {
	return &dto.ProductDTO{
		ID:          product.ID.String(),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		SKU:         product.SKU,
		Categories:  categoryMapper.ToCategoryDTOs(product.Categories),
		UOMs:        uomMapper.ToUOMDTOs(product.UOMs),
	}
}

func ToProductDTOs(products []*domain.Product) []*dto.ProductDTO {
	var productDTOs []*dto.ProductDTO
	for _, product := range products {
		productDTOs = append(productDTOs, &dto.ProductDTO{
			ID:          product.ID.String(),
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			SKU:         product.SKU,
			Categories:  categoryMapper.ToCategoryDTOs(product.Categories),
			UOMs:        uomMapper.ToUOMDTOs(product.UOMs),
		})
	}
	return productDTOs
}
