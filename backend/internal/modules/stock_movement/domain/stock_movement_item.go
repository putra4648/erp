package domain

import (
	productDomain "putra4648/erp/internal/modules/product/domain"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type StockMovementItem struct {
	ID              uuid.UUID             `gorm:"type:uuid;primarykey;default:uuid_generate_v4()"`
	StockMovementID uuid.UUID             `gorm:"type:uuid;column:stock_movement_id;not null"`
	ProductID       uuid.UUID             `gorm:"type:uuid;column:product_id;not null"`
	Product         productDomain.Product `gorm:"foreignKey:ProductID"`
	Quantity        decimal.Decimal       `gorm:"type:decimal(19,4);not null"`
	Note            string                `gorm:"type:text"`
}

type StockMovementItemResponse struct {
	ID        uuid.UUID                     `json:"id"`
	ProductID uuid.UUID                     `json:"product_id"`
	Product   productDomain.ProductResponse `json:"product"`
	Quantity  decimal.Decimal               `json:"quantity"`
	Note      string                        `json:"note"`
}

type StockMovementItemDTO struct {
	ID        string          `json:"id"`
	ProductID string          `json:"product_id" validate:"required"`
	Quantity  decimal.Decimal `json:"quantity" validate:"required,gt=0"`
	Note      string          `json:"note"`
}

func (i *StockMovementItem) ToResponse() *StockMovementItemResponse {
	return &StockMovementItemResponse{
		ID:        i.ID,
		ProductID: i.ProductID,
		Product:   *i.Product.ToResponse(),
		Quantity:  i.Quantity,
		Note:      i.Note,
	}
}
