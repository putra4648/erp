package service

import (
	"context"
	"putra4648/erp/internal/modules/inventory/warehouse/domain"
	warehouseDto "putra4648/erp/internal/modules/inventory/warehouse/dto"
	"putra4648/erp/internal/modules/inventory/warehouse/repository"

	"putra4648/erp/internal/modules/shared/dto"

	"github.com/google/uuid"
)

type WarehouseQueryService interface {
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Warehouse, error)
	FindAll(ctx context.Context, req *warehouseDto.WarehouseFindAllRequest) (*dto.PaginationResponse[*domain.Warehouse], error)
}

type warehouseQueryService struct {
	repo repository.WarehouseRepository
}

func NewWarehouseQueryService(repo repository.WarehouseRepository) WarehouseQueryService {
	return &warehouseQueryService{repo: repo}
}

func (s *warehouseQueryService) FindByID(ctx context.Context, id uuid.UUID) (*domain.Warehouse, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *warehouseQueryService) FindAll(ctx context.Context, req *warehouseDto.WarehouseFindAllRequest) (*dto.PaginationResponse[*domain.Warehouse], error) {
	warehouses, total, err := s.repo.FindAll(ctx, req)
	if err != nil {
		return nil, err
	}

	return &dto.PaginationResponse[*domain.Warehouse]{
		Items: warehouses,
		Total: total,
		Page:  req.Page,
		Size:  req.Size,
	}, nil
}
