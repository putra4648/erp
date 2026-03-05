package mapper

import (
	"putra4648/erp/internal/modules/stock_adjustment/domain"
	"putra4648/erp/internal/modules/stock_adjustment/dto"
)

func ToStockAdjustmentDto(stockAdjustment *domain.StockAdjustment) *dto.StockAdjustmentDto {
	items := make([]dto.StockAdjustmentItemDto, len(stockAdjustment.Items))
	for i, item := range stockAdjustment.Items {
		items[i] = *ToStockAdjustmentItemDto(&item)
	}
	return &dto.StockAdjustmentDto{
		ID:              stockAdjustment.ID,
		AdjustmentNo:    stockAdjustment.AdjustmentNo,
		WarehouseID:     stockAdjustment.WarehouseID,
		TransactionDate: stockAdjustment.TransactionDate,
		Status:          string(stockAdjustment.Status),
		Note:            stockAdjustment.Note,
		CreatedBy:       stockAdjustment.CreatedBy,
		ApprovedBy:      stockAdjustment.ApprovedBy,
		Items:           items,
	}
}

func ToStockAdjustmentItemDto(stockAdjustmentItem *domain.StockAdjustmentItem) *dto.StockAdjustmentItemDto {
	return &dto.StockAdjustmentItemDto{
		ID:            stockAdjustmentItem.ID,
		ProductID:     stockAdjustmentItem.ProductID,
		ProductName:   stockAdjustmentItem.Product.Name,
		ReasonID:      stockAdjustmentItem.ReasonID,
		ReasonName:    stockAdjustmentItem.Reason.Name,
		ActualQty:     stockAdjustmentItem.ActualQty,
		SystemQty:     stockAdjustmentItem.SystemQty,
		AdjustmentQty: stockAdjustmentItem.AdjustmentQty,
	}
}
