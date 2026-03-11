package domain

import (
	"github.com/google/uuid"
)

type Category struct {
	ID   *uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name string     `gorm:"not null;size:255"`
}
