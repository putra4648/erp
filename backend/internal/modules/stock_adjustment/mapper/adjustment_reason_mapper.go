package mapper

import (
	"putra4648/erp/internal/modules/stock_adjustment/domain"
	"putra4648/erp/internal/modules/stock_adjustment/dto"
)

func ToAdjustmentReasonDto(reason *domain.AdjustmentReason) *dto.AdjustmentReasonDto {
	if reason == nil {
		return nil
	}
	return &dto.AdjustmentReasonDto{
		ID:   reason.ID.String(),
		Name: reason.Name,
		Code: reason.AccountCode,
	}
}
