package domain

import (
	categoryDomain "putra4648/erp/internal/category/domain"
	supplierDomain "putra4648/erp/internal/supplier/domain"
	uomDomain "putra4648/erp/internal/uom/domain"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Product struct {
	ID   uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name string    `gorm:"not null;size:255"`
	// Description string                     `gorm:"type:text"`
	SKU        string                  `gorm:"unique;not null;size:100"`
	Price      decimal.Decimal         `gorm:"not null;precision:19;scale:2"`
	SupplierID uuid.UUID               `gorm:"type:uuid;column:supplier_id;not null"`
	Supplier   supplierDomain.Supplier `gorm:"foreignKey:SupplierID"`
	// Quantity   int                        `gorm:"not null;default:0"`
	Categories []*categoryDomain.Category `gorm:"many2many:product_categories;"`
	UOMs       []*uomDomain.UOM           `gorm:"many2many:product_uoms;"`
	IsActive   bool                       `gorm:"not null;default:true"`
}
