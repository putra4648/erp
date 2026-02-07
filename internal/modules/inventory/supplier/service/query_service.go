package service

import (
	"context"
	"putra4648/erp/internal/modules/inventory/supplier/domain"
	supplierDto "putra4648/erp/internal/modules/inventory/supplier/dto"
	"putra4648/erp/internal/modules/inventory/supplier/repository"

	"putra4648/erp/internal/modules/shared/dto"

	"github.com/google/uuid"
)

type SupplierQueryService interface {
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Supplier, error)
	FindAll(ctx context.Context, req *supplierDto.SupplierFindAllRequest) (*dto.PaginationResponse[*domain.Supplier], error)
}

type supplierQueryService struct {
	repo repository.SupplierRepository
}

func NewSupplierQueryService(repo repository.SupplierRepository) SupplierQueryService {
	return &supplierQueryService{repo: repo}
}

func (s *supplierQueryService) FindByID(ctx context.Context, id uuid.UUID) (*domain.Supplier, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *supplierQueryService) FindAll(ctx context.Context, req *supplierDto.SupplierFindAllRequest) (*dto.PaginationResponse[*domain.Supplier], error) {
	suppliers, total, err := s.repo.FindAll(ctx, req)
	if err != nil {
		return nil, err
	}

	return &dto.PaginationResponse[*domain.Supplier]{
		Items: suppliers,
		Total: total,
		Page:  req.Page,
		Size:  req.Size,
	}, nil
}
