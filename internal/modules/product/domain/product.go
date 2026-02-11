package domain

import (
	. "putra4648/erp/internal/modules/shared/utils"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Product struct {
	ID          uuid.UUID       `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string          `gorm:"not null;size:255"`
	Description string          `gorm:"type:text"`
	SKU         string          `gorm:"unique;not null;size:100"`
	Price       decimal.Decimal `gorm:"not null;precision:19;scale:2"`
	Cost        decimal.Decimal `gorm:"not null;precision:19;scale:2"`
	Quantity    int             `gorm:"not null;default:0"`
	Categories  []*Category     `gorm:"many2many:product_categories;"`
	UOMs        []*UOM          `gorm:"many2many:product_uoms;"`
	IsActive    bool            `gorm:"not null;default:true"`
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
		Categories:  MapSlice(p.Categories, func(c *Category) *CategoryResponse { return c.ToResponse() }),
		UOMs:        MapSlice(p.UOMs, func(u *UOM) *UOMResponse { return u.ToResponse() }),
		IsActive:    p.IsActive,
	}
}

type ProductResponse struct {
	ID          uuid.UUID           `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	SKU         string              `json:"sku"`
	Price       decimal.Decimal     `json:"price"`
	Cost        decimal.Decimal     `json:"cost"`
	Quantity    int                 `json:"quantity"`
	Categories  []*CategoryResponse `json:"categories"`
	UOMs        []*UOMResponse      `json:"uoms"`
	IsActive    bool                `json:"is_active"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
}

type ProductDTO struct {
	ID          string          `json:"id"`
	Name        string          `json:"name" validate:"required,max=255"`
	Description string          `json:"description" validate:"max:65000"`
	SKU         string          `json:"sku" validate:"required,max=100,alphanum"`
	Price       decimal.Decimal `json:"price" validate:"required,gt=0"`
	Cost        decimal.Decimal `json:"cost" validate:"required,gte=0"`
	Quantity    int             `json:"quantity" validate:"gte=0"`
	Categories  []CategoryDTO   `json:"categories"`
	UOMs        []UOMDTO        `json:"uoms"`
	IsActive    bool            `json:"is_active" default:"true"`
}

func (dto *ProductDTO) ToModel() *Product {
	id, _ := uuid.Parse(dto.ID) // Returns uuid.Nil if invalid/empty

	categories := make([]*Category, len(dto.Categories))
	for i, cat := range dto.Categories {
		categories[i] = &Category{ID: uuid.MustParse(cat.ID)}
	}
	uoms := make([]*UOM, len(dto.UOMs))
	for i, uom := range dto.UOMs {
		uoms[i] = &UOM{ID: uuid.MustParse(uom.ID)}
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
