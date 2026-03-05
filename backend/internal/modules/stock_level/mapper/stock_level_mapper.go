package mapper

import (
	"putra4648/erp/internal/modules/stock_level/domain"
	"putra4648/erp/internal/modules/stock_level/dto"
	"time"
)

func ToStockLevelResponse(model *domain.StockLevel) *dto.StockLevelResponse {
	return &dto.StockLevelResponse{
		ID:            model.ID.String(),
		ProductID:     model.ProductID.String(),
		ProductName:   model.Product.Name,
		WarehouseID:   model.WarehouseID.String(),
		WarehouseName: model.Warehouse.Name,
		Quantity:      model.Quantity,
		LastUpdated:   model.LastUpdated.Format(time.RFC3339),
	}
}

func ToStockLevelResponses(models []*domain.StockLevel) []*dto.StockLevelResponse {
	responses := make([]*dto.StockLevelResponse, len(models))
	for i, m := range models {
		responses[i] = ToStockLevelResponse(m)
	}
	return responses
}
