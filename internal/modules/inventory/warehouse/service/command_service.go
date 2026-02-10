package service

import (
	"context"
	"putra4648/erp/internal/modules/inventory/warehouse/domain"
	"putra4648/erp/internal/modules/inventory/warehouse/dto"
	"putra4648/erp/internal/modules/inventory/warehouse/repository"

	"github.com/google/uuid"
)

type WarehouseCommandService interface {
	Create(ctx context.Context, req *dto.CreateWarehouseRequest) (uuid.UUID, error)
	Update(ctx context.Context, id uuid.UUID, req *dto.UpdateWarehouseRequest) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type warehouseCommandService struct {
	repo repository.WarehouseRepository
}

func NewWarehouseCommandService(repo repository.WarehouseRepository) WarehouseCommandService {
	return &warehouseCommandService{repo: repo}
}

func (s *warehouseCommandService) Create(ctx context.Context, req *dto.CreateWarehouseRequest) (uuid.UUID, error) {
	warehouse := &domain.Warehouse{
		ID:   uuid.New(),
		Name: req.Name,
		Code: req.Code,
	}
	if err := s.repo.Save(ctx, warehouse); err != nil {
		return uuid.Nil, err
	}
	return warehouse.ID, nil
}

func (s *warehouseCommandService) Update(ctx context.Context, id uuid.UUID, req *dto.UpdateWarehouseRequest) error {
	warehouse, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	warehouse.Name = req.Name
	return s.repo.Update(ctx, warehouse)
}

func (s *warehouseCommandService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
