package domain

import (
	"github.com/google/uuid"
)

type UOM struct {
	ID   uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name string    `gorm:"not null;size:255"`
}
