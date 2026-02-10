package domain

import (
	uomDomain "putra4648/erp/internal/modules/uom/domain"

	"github.com/google/uuid"
)

type ProductUOM struct {
	ProductID uuid.UUID     `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UOMID     uuid.UUID     `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UOM       uomDomain.UOM `gorm:"foreignKey:UOMID"`
}

type ProductUOMResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (pu *ProductUOM) ToResponse() *ProductUOMResponse {
	return &ProductUOMResponse{
		ID:   pu.UOMID,
		Name: pu.UOM.Name,
	}
}

type ProductUOMDTO struct {
	ID uuid.UUID `json:"id"`
}

func (pu *ProductUOM) ToDTO() *ProductUOMDTO {
	return &ProductUOMDTO{
		ID: pu.UOMID,
	}
}

func (pu *ProductUOMDTO) ToModel() *ProductUOM {
	return &ProductUOM{
		UOMID: pu.ID,
	}
}
