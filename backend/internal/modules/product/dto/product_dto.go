package dto

import (
	categoryDto "putra4648/erp/internal/modules/category/dto"
	uomDto "putra4648/erp/internal/modules/uom/dto"

	"github.com/shopspring/decimal"
)

type ProductRequest struct {
	Name string `json:"name"`
	Page int    `json:"page"`
	Size int    `json:"size"`
}

type ProductDTO struct {
	ID          string                     `json:"id"`
	Name        string                     `json:"name" validate:"required,max=255"`
	Description string                     `json:"description" validate:"max:65000"`
	SKU         string                     `json:"sku" validate:"required,max=100,alphanum"`
	Price       decimal.Decimal            `json:"price" validate:"required,gt=0"`
	Cost        decimal.Decimal            `json:"cost" validate:"required,gte=0"`
	Quantity    int                        `json:"quantity" validate:"gte=0"`
	Categories  []*categoryDto.CategoryDTO `json:"categories"`
	UOMs        []*uomDto.UOMDTO           `json:"uoms"`
	SupplierID  string                     `json:"supplier_id" validate:"required"`
	IsActive    bool                       `json:"is_active" default:"true"`
}
