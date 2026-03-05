package service

import (
	"context"
	"putra4648/erp/internal/modules/shared/dto"
	warehouseDto "putra4648/erp/internal/modules/warehouse/dto"
	"putra4648/erp/internal/modules/warehouse/mapper"
	"putra4648/erp/internal/modules/warehouse/repository"

	"github.com/google/uuid"
)

type warehouseQueryService struct {
	repo repository.WarehouseRepository
}

func NewWarehouseQueryService(repo repository.WarehouseRepository) WarehouseQueryService {
	return &warehouseQueryService{repo: repo}
}

func (s *warehouseQueryService) FindByID(ctx context.Context, id uuid.UUID) (*warehouseDto.WarehouseDto, error) {
	warehouse, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, &WarehouseError{Code: "NOT_FOUND", Message: "Warehouse not found"}
	}
	return mapper.ToWarehouseDto(warehouse), nil
}

func (s *warehouseQueryService) FindAll(ctx context.Context, req *warehouseDto.WarehouseFindAllRequest) (*dto.PaginationResponse[*warehouseDto.WarehouseDto], error) {
	warehouses, total, err := s.repo.FindAll(ctx, req)
	if err != nil {
		return nil, err
	}

	warehouseDTOs := mapper.ToWarehouseDtos(warehouses)

	return &dto.PaginationResponse[*warehouseDto.WarehouseDto]{
		Items: warehouseDTOs,
		Total: total,
		Page:  req.Page,
		Size:  req.Size,
	}, nil
}
