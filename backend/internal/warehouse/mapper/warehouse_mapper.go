package mapper

import (
	"putra4648/erp/internal/warehouse/domain"
	"putra4648/erp/internal/warehouse/dto"
)

func ToWarehouse(req *dto.WarehouseDTO) *domain.Warehouse {
	return &domain.Warehouse{
		Name: req.Name,
		Code: req.Code,
	}
}

func ToWarehouseDto(warehouse *domain.Warehouse) *dto.WarehouseDTO {
	return &dto.WarehouseDTO{
		ID:   warehouse.ID.String(),
		Name: warehouse.Name,
		Code: warehouse.Code,
	}
}

func ToWarehouseDtos(warehouses []*domain.Warehouse) []*dto.WarehouseDTO {
	warehouseDtos := make([]*dto.WarehouseDTO, len(warehouses))
	for i, warehouse := range warehouses {
		warehouseDtos[i] = ToWarehouseDto(warehouse)
	}
	return warehouseDtos
}
