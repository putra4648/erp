package model

import (
	categoryModel "putra4648/erp/internal/modules/category/model"
	uomModel "putra4648/erp/internal/modules/uom/model"
	. "putra4648/erp/utils"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	*gorm.Model
	ID          uuid.UUID       `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string          `gorm:"not null;size:255"`
	Description string          `gorm:"type:text"`
	SKU         string          `gorm:"unique;not null;size:100"`
	Price       decimal.Decimal `gorm:"not null;precision:19;scale:2"`
	Cost        decimal.Decimal `gorm:"not null;precision:19;scale:2"`
	Quantity    int             `gorm:"not null;default:0"`
	Categories  []*categoryModel.Category     `gorm:"many2many:product_categories;"`
	UOMs        []*uomModel.UOM            `gorm:"many2many:product_uoms;"`
	IsActive    bool            `gorm:"not null;default:true"`
}

type ProductDTO struct {
	Name        string          `json:"name" validate:"required,max=255"`
	Description string          `json:"description" validate:"max:65000"`
	SKU         string          `json:"sku" validate:"required,max=100,alphanum"`
	Price       decimal.Decimal `json:"price" validate:"required,gt=0"`
	Cost        decimal.Decimal `json:"cost" validate:"required,gte=0"`
	Quantity    int             `json:"quantity" validate:"gte=0"`
	CategoryIDs []uuid.UUID     `json:"category_ids"`
	UOMIDs      []uuid.UUID       `json:"uom_ids"`
	IsActive    bool            `json:"is_active" default:"true"`
}

type ProductResponse struct {
	ID          uuid.UUID           `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	SKU         string              `json:"sku"`
	Price       decimal.Decimal     `json:"price"`
	Cost        decimal.Decimal     `json:"cost"`
	Quantity    int                 `json:"quantity"`
	Categories  []*categoryModel.CategoryResponse `json:"categories"`
	UOMs        []*uomModel.UOMResponse        `json:"uoms"`
	IsActive    bool                `json:"is_active"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
}

func (p *Product) ToResponse() *ProductResponse {
	return &ProductResponse{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		SKU:         p.SKU,
		Price:       p.Price,
		Cost:        p.Cost,
		Quantity:    p.Quantity,
		Categories:  MapSlice(p.Categories, func(c *categoryModel.Category) *categoryModel.CategoryResponse { return c.ToResponse() }),
		UOMs:        MapSlice(p.UOMs, func(u *uomModel.UOM) *uomModel.UOMResponse { return u.ToResponse() }),
		IsActive:    p.IsActive,
		CreatedAt:   p.CreatedAt.UTC(),
		UpdatedAt:   p.UpdatedAt.UTC(),
	}
}

func (dto *ProductDTO) ToModel() *Product {
	categories := make([]*categoryModel.Category, len(dto.CategoryIDs))
	for i, catID := range dto.CategoryIDs {
		categories[i] = &categoryModel.Category{ID: catID}
	}
	uoms := make([]*uomModel.UOM, len(dto.UOMIDs))
	for i, uomID := range dto.UOMIDs {
		uoms[i] = &uomModel.UOM{ID: uomID}
	}

	return &Product{
		Name:        dto.Name,
		Description: dto.Description,
		SKU:         dto.SKU,
		Price:       dto.Price,
		Cost:        dto.Cost,
		Quantity:    dto.Quantity,
		Categories:  categories,
		UOMs:        uoms,
		IsActive:    dto.IsActive,
	}
}
