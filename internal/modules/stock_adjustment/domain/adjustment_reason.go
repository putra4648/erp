package domain

import (
	"putra4648/erp/internal/modules/stock_adjustment/dto"

	"github.com/google/uuid"
)

type AdjustmentReason struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	AccountCode string    `gorm:"type:varchar(20);not null" json:"account_code"`
}

func (r *AdjustmentReason) ToResponse() *dto.StockAdjustmentItemResponse {
	// This might not be right since AdjustmentReason is not a full item, but let's see.
	// Actually, let's just add a generic ToResponse if needed.
	return nil
}
