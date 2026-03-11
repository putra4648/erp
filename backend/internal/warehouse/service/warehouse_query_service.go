package service

import (
	"context"
	"putra4648/erp/internal/shared/dto"
	"putra4648/erp/internal/shared/errors"
	warehouseDto "putra4648/erp/internal/warehouse/dto"
	"putra4648/erp/internal/warehouse/mapper"
	"putra4648/erp/internal/warehouse/repository"

	"github.com/google/uuid"
)

type warehouseQueryService struct {
	repo repository.WarehouseRepository
}

func NewWarehouseQueryService(repo repository.WarehouseRepository) WarehouseQueryService {
	return &warehouseQueryService{repo: repo}
}

func (s *warehouseQueryService) FindByID(ctx context.Context, id uuid.UUID) (*warehouseDto.WarehouseDTO, error) {
	warehouse, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, &errors.ErrorDto{Code: "NOT_FOUND", Message: "Warehouse not found"}
	}
	return mapper.ToWarehouseDto(warehouse), nil
}

func (s *warehouseQueryService) FindAll(ctx context.Context, pagination *dto.PaginationRequest, req *warehouseDto.WarehouseDTO) (*dto.PaginationResponse[*warehouseDto.WarehouseDTO], error) {
	warehouses, total, err := s.repo.FindAll(ctx, pagination, req)
	if err != nil {
		return nil, err
	}

	warehouseDTOs := mapper.ToWarehouseDtos(warehouses)

	return &dto.PaginationResponse[*warehouseDto.WarehouseDTO]{
		Items: warehouseDTOs,
		Total: total,
		Page:  pagination.Page,
		Size:  pagination.Size,
	}, nil
}
