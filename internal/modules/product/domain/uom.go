package domain

import "github.com/google/uuid"

type UOM struct {
	ID   uuid.UUID `gorm:"type:uuid;primarykey;default:uuid_generate_v4()"`
	Name string    `gorm:"type:varchar(255)"`
}

type UOMResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type UOMDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (dto *UOMDTO) ToModel() *UOM {
	id, _ := uuid.Parse(dto.ID)
	return &UOM{
		ID:   id,
		Name: dto.Name,
	}
}

func (u *UOM) ToResponse() *UOMResponse {
	return &UOMResponse{
		ID:   u.ID,
		Name: u.Name,
	}
}
