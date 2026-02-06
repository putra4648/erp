package model

import (
	. "putra4648/erp/utils"
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
	Category    []Category      `gorm:"foreignKey:ProductID"`
	UomID       uuid.UUID       `gorm:"type:uuid;not null"`
	UOM         UOM             `gorm:"foreignKey:UomID"`
	IsActive    bool            `gorm:"not null;default:true"`
	CreatedAt   time.Time       `gorm:"not null"`
	UpdatedAt   time.Time       `gorm:"not null"`
	DeletedAt   *time.Time      `gorm:"nullable"`
}

type ProductDTO struct {
	Name        string          `json:"name" validate:"required,max=255"`
	Description string          `json:"description" validate:"max:65000"`
	SKU         string          `json:"sku" validate:"required,max=100,alphanum"`
	Price       decimal.Decimal `json:"price" validate:"required,gt=0"`
	Cost        decimal.Decimal `json:"cost" validate:"required,gte=0"`
	Quantity    int             `json:"quantity" validate:"gte=0"`
	Category    []CategoryDTO   `json:"category"`
	UOM         UOMDTO          `json:"uom"`
	IsActive    bool            `json:"is_active" default:"true"`
}

type ProductResponse struct {
	ID          uuid.UUID          `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	SKU         string             `json:"sku"`
	Price       decimal.Decimal    `json:"price"`
	Cost        decimal.Decimal    `json:"cost"`
	Quantity    int                `json:"quantity"`
	Category    []CategoryResponse `json:"category"`
	UOM         UOMResponse        `json:"uom"`
	IsActive    bool               `json:"is_active"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
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
		Category:    MapSlice(p.Category, func(c Category) CategoryResponse { return *c.ToResponse() }),
		UOM:         *p.UOM.ToResponse(),
		IsActive:    p.IsActive,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func (dto *ProductDTO) ToModel() *Product {

	return &Product{
		Name:        dto.Name,
		Description: dto.Description,
		SKU:         dto.SKU,
		Price:       dto.Price,
		Cost:        dto.Cost,
		Quantity:    dto.Quantity,
		Category:    MapSlice(dto.Category, func(c CategoryDTO) Category { return *c.ToModel() }),
		UOM:         *dto.UOM.ToModel(),
		IsActive:    dto.IsActive,
	}
}
