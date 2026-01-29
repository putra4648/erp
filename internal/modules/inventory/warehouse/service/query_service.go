package service

import (
	"context"
	"putra4648/erp/internal/modules/inventory/warehouse/domain"

	"github.com/google/uuid"
	"putra4648/erp/internal/modules/shared/dto"
)

type WarehouseQueryService interface {
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Warehouse, error)
	FindAll(ctx context.Context, page, size int) (*dto.PaginationResponse[*domain.Warehouse], error)
}

type warehouseQueryService struct {
	repo domain.WarehouseRepository
}

func NewWarehouseQueryService(repo domain.WarehouseRepository) WarehouseQueryService {
	return &warehouseQueryService{repo: repo}
}

func (s *warehouseQueryService) FindByID(ctx context.Context, id uuid.UUID) (*domain.Warehouse, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *warehouseQueryService) FindAll(ctx context.Context, page, size int) (*dto.PaginationResponse[*domain.Warehouse], error) {
	warehouses, total, err := s.repo.FindAll(ctx, page, size)
	if err != nil {
		return nil, err
	}

	return &dto.PaginationResponse[*domain.Warehouse]{
		Items: warehouses,
		Total: total,
		Page:  page,
		Size:  size,
	}, nil
}
