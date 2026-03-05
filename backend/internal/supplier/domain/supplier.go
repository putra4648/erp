package domain

import "github.com/google/uuid"

type Supplier struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Name    string    `gorm:"type:varchar(100);not null" json:"name"`
	Code    string    `gorm:"type:varchar(20);not null;unique" json:"code"`
	Address string    `gorm:"type:text" json:"address"`
	Phone   string    `gorm:"type:varchar(20)" json:"phone"`
	Email   string    `gorm:"type:varchar(100)" json:"email"`
}
