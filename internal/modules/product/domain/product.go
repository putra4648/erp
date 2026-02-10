package domain

import (
	. "putra4648/erp/utils"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Product struct {
	ID          uuid.UUID          `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string             `gorm:"not null;size:255"`
	Description string             `gorm:"type:text"`
	SKU         string             `gorm:"unique;not null;size:100"`
	Price       decimal.Decimal    `gorm:"not null;precision:19;scale:2"`
	Cost        decimal.Decimal    `gorm:"not null;precision:19;scale:2"`
	Quantity    int                `gorm:"not null;default:0"`
	Categories  []*ProductCategory `gorm:"foreignKey:ProductID"`
	UOMs        []*ProductUOM      `gorm:"foreignKey:ProductID"`
	IsActive    bool               `gorm:"not null;default:true"`
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
		Categories:  MapSlice(p.Categories, func(c *ProductCategory) *ProductCategoryResponse { return c.ToResponse() }),
		UOMs:        MapSlice(p.UOMs, func(u *ProductUOM) *ProductUOMResponse { return u.ToResponse() }),
		IsActive:    p.IsActive,
	}
}

type ProductResponse struct {
	ID          uuid.UUID                  `json:"id"`
	Name        string                     `json:"name"`
	Description string                     `json:"description"`
	SKU         string                     `json:"sku"`
	Price       decimal.Decimal            `json:"price"`
	Cost        decimal.Decimal            `json:"cost"`
	Quantity    int                        `json:"quantity"`
	Categories  []*ProductCategoryResponse `json:"categories"`
	UOMs        []*ProductUOMResponse      `json:"uoms"`
	IsActive    bool                       `json:"is_active"`
	CreatedAt   time.Time                  `json:"created_at"`
	UpdatedAt   time.Time                  `json:"updated_at"`
}

type ProductDTO struct {
	ID          string               `json:"id"`
	Name        string               `json:"name" validate:"required,max=255"`
	Description string               `json:"description" validate:"max:65000"`
	SKU         string               `json:"sku" validate:"required,max=100,alphanum"`
	Price       decimal.Decimal      `json:"price" validate:"required,gt=0"`
	Cost        decimal.Decimal      `json:"cost" validate:"required,gte=0"`
	Quantity    int                  `json:"quantity" validate:"gte=0"`
	Categories  []ProductCategoryDTO `json:"categories"`
	UOMs        []ProductUOMDTO      `json:"uoms"`
	IsActive    bool                 `json:"is_active" default:"true"`
}

func (dto *ProductDTO) ToModel() *Product {
	id, _ := uuid.Parse(dto.ID) // Returns uuid.Nil if invalid/empty

	categories := make([]*ProductCategory, len(dto.Categories))
	for i, cat := range dto.Categories {
		categories[i] = &ProductCategory{CategoryID: cat.ID, ProductID: id}
	}
	uoms := make([]*ProductUOM, len(dto.UOMs))
	for i, uom := range dto.UOMs {
		uoms[i] = &ProductUOM{UOMID: uom.ID, ProductID: id}
	}

	return &Product{
		ID:          id,
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
