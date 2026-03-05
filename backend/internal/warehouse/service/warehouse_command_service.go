package service

import (
	"context"
	"putra4648/erp/internal/warehouse/domain"
	"putra4648/erp/internal/warehouse/dto"
	"putra4648/erp/internal/warehouse/mapper"
	"putra4648/erp/internal/warehouse/repository"

	"github.com/google/uuid"
)

type warehouseCommandService struct {
	repo repository.WarehouseRepository
}

func NewWarehouseCommandService(repo repository.WarehouseRepository) WarehouseCommandService {
	return &warehouseCommandService{repo: repo}
}

func (s *warehouseCommandService) Create(ctx context.Context, req *dto.WarehouseDto) (*dto.WarehouseDto, error) {
	warehouse := &domain.Warehouse{
		ID:   uuid.New(),
		Name: req.Name,
		Code: req.Code,
	}

	if err := s.repo.Save(ctx, warehouse); err != nil {
		return nil, &WarehouseError{Code: "DATABASE_ERROR", Message: "Failed to create Warehouse"}
	}
	return mapper.ToWarehouseDto(warehouse), nil
}

func (s *warehouseCommandService) Update(ctx context.Context, req *dto.WarehouseDto) (*dto.WarehouseDto, error) {
	id, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, &WarehouseError{Code: "INVALID_ID", Message: "Invalid Warehouse ID"}
	}
	warehouse, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, &WarehouseError{Code: "NOT_FOUND", Message: "Warehouse not found"}
	}
	warehouse.Name = req.Name
	if err := s.repo.Update(ctx, warehouse); err != nil {
		return nil, &WarehouseError{Code: "DATABASE_ERROR", Message: "Failed to update Warehouse"}
	}
	return mapper.ToWarehouseDto(warehouse), nil
}

func (s *warehouseCommandService) Delete(ctx context.Context, id uuid.UUID) (*dto.WarehouseDto, error) {
	warehouse, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, &WarehouseError{Code: "NOT_FOUND", Message: "Warehouse not found"}
	}
	if err := s.repo.Delete(ctx, id); err != nil {
		return nil, &WarehouseError{Code: "DATABASE_ERROR", Message: "Failed to delete Warehouse"}
	}
	return mapper.ToWarehouseDto(warehouse), nil
}
