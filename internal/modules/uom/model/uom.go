package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UOM struct {
	*gorm.Model
	ID   uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name string    `gorm:"not null;size:255"`
}

type UOMDTO struct {
	Name string `json:"name" validate:"required,max=255"`
}

type UOMResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (u *UOM) ToResponse() *UOMResponse {
	return &UOMResponse{
		ID:   u.ID,
		Name: u.Name,
	}
}

func (dto *UOMDTO) ToModel() *UOM {
	return &UOM{
		Name: dto.Name,
	}
}
