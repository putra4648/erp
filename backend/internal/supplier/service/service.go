package service

import (
	"context"
	sharedDto "putra4648/erp/internal/shared/dto"
	"putra4648/erp/internal/supplier/dto"
)

type SupplierCommandService interface {
	Create(ctx context.Context, req *dto.SupplierDTO) (*dto.SupplierDTO, error)
	Update(ctx context.Context, req *dto.SupplierDTO) (*dto.SupplierDTO, error)
	Delete(ctx context.Context, id string) (*dto.SupplierDTO, error)
}

type SupplierQueryService interface {
	FindByID(ctx context.Context, id string) (*dto.SupplierDTO, error)
	FindAll(ctx context.Context, pagination *sharedDto.PaginationRequest, req *dto.SupplierDTO) (*sharedDto.PaginationResponse[*dto.SupplierDTO], error)
}
