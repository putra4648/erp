package service

import (
	"context"
	sharedDto "putra4648/erp/internal/modules/shared/dto"
	"putra4648/erp/internal/modules/warehouse/dto"
	warehouseDto "putra4648/erp/internal/modules/warehouse/dto"

	"github.com/google/uuid"
)

type WarehouseCommandService interface {
	Create(ctx context.Context, req *dto.WarehouseDto) (*dto.WarehouseDto, error)
	Update(ctx context.Context, req *dto.WarehouseDto) (*dto.WarehouseDto, error)
	Delete(ctx context.Context, id uuid.UUID) (*dto.WarehouseDto, error)
}

type WarehouseQueryService interface {
	FindByID(ctx context.Context, id uuid.UUID) (*dto.WarehouseDto, error)
	FindAll(ctx context.Context, req *warehouseDto.WarehouseFindAllRequest) (*sharedDto.PaginationResponse[*dto.WarehouseDto], error)
}
