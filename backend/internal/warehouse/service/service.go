package service

import (
	"context"
	sharedDto "putra4648/erp/internal/shared/dto"
	"putra4648/erp/internal/warehouse/dto"

	"github.com/google/uuid"
)

type WarehouseCommandService interface {
	Create(ctx context.Context, req *dto.WarehouseDTO) (*dto.WarehouseDTO, error)
	Update(ctx context.Context, req *dto.WarehouseDTO) (*dto.WarehouseDTO, error)
	Delete(ctx context.Context, id uuid.UUID) (*dto.WarehouseDTO, error)
}

type WarehouseQueryService interface {
	FindByID(ctx context.Context, id uuid.UUID) (*dto.WarehouseDTO, error)
	FindAll(ctx context.Context, pagination *sharedDto.PaginationRequest, req *dto.WarehouseDTO) (*sharedDto.PaginationResponse[*dto.WarehouseDTO], error)
}
