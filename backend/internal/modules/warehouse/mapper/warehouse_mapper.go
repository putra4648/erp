package mapper

import (
	"putra4648/erp/internal/modules/warehouse/domain"
	"putra4648/erp/internal/modules/warehouse/dto"
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
	var warehouseDtos []*dto.WarehouseDto
	for _, warehouse := range warehouses {
		warehouseDtos = append(warehouseDtos, ToWarehouseDto(warehouse))
	}
	return warehouseDtos
}
