package mapper

import (
	"putra4648/erp/internal/stock_level/domain"
	"putra4648/erp/internal/stock_level/dto"
	"time"
)

func ToStockLevelResponse(model *domain.StockLevel) *dto.StockLevelDto {
	productId := model.ProductID.String()
	warehouseId := model.WarehouseID.String()
	return &dto.StockLevelDto{
		ID:            model.ID.String(),
		ProductID:     &productId,
		ProductName:   model.Product.Name,
		WarehouseID:   &warehouseId,
		WarehouseName: model.Warehouse.Name,
		Quantity:      model.Quantity,
		LastUpdated:   model.LastUpdated.Format(time.RFC3339),
	}
}

func ToStockLevelResponses(models []*domain.StockLevel) []*dto.StockLevelDto {
	responses := make([]*dto.StockLevelDto, len(models))
	for i, m := range models {
		responses[i] = ToStockLevelResponse(m)
	}
	return responses
}
