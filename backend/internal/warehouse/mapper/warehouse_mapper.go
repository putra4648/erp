package mapper

import (
	"putra4648/erp/internal/warehouse/domain"
	"putra4648/erp/internal/warehouse/dto"
)

func ToWarehouse(req *dto.WarehouseDto) *domain.Warehouse {
	return &domain.Warehouse{
		Name: req.Name,
		Code: req.Code,
	}
}

func ToWarehouseDto(warehouse *domain.Warehouse) *dto.WarehouseDto {
	return &dto.WarehouseDto{
		ID:   warehouse.ID.String(),
		Name: warehouse.Name,
		Code: warehouse.Code,
	}
}

func ToWarehouseDtos(warehouses []*domain.Warehouse) []*dto.WarehouseDto {
	warehouseDtos := make([]*dto.WarehouseDto, len(warehouses))
	for i, warehouse := range warehouses {
		warehouseDtos[i] = ToWarehouseDto(warehouse)
	}
	return warehouseDtos
}
