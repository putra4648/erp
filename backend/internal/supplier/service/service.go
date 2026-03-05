package service

import (
	"context"
	sharedDto "putra4648/erp/internal/shared/dto"
	"putra4648/erp/internal/supplier/dto"
)

type SupplierCommandService interface {
	Create(ctx context.Context, req *dto.SupplierDto) (*dto.SupplierDto, error)
	Update(ctx context.Context, req *dto.SupplierDto) (*dto.SupplierDto, error)
	Delete(ctx context.Context, id string) (*dto.SupplierDto, error)
}

type SupplierQueryService interface {
	FindByID(ctx context.Context, id string) (*dto.SupplierDto, error)
	FindAll(ctx context.Context, req *dto.SupplierFindAllRequest) (*sharedDto.PaginationResponse[*dto.SupplierDto], error)
}
