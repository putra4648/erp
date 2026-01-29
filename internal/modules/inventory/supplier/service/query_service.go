package service

import (
	"context"
	"putra4648/erp/internal/modules/inventory/supplier/domain"

	"github.com/google/uuid"
	"putra4648/erp/internal/modules/shared/dto"
)

type SupplierQueryService interface {
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Supplier, error)
	FindAll(ctx context.Context, page, size int) (*dto.PaginationResponse[*domain.Supplier], error)
}

type supplierQueryService struct {
	repo domain.SupplierRepository
}

func NewSupplierQueryService(repo domain.SupplierRepository) SupplierQueryService {
	return &supplierQueryService{repo: repo}
}

func (s *supplierQueryService) FindByID(ctx context.Context, id uuid.UUID) (*domain.Supplier, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *supplierQueryService) FindAll(ctx context.Context, page, size int) (*dto.PaginationResponse[*domain.Supplier], error) {
	suppliers, total, err := s.repo.FindAll(ctx, page, size)
	if err != nil {
		return nil, err
	}

	return &dto.PaginationResponse[*domain.Supplier]{
		Items: suppliers,
		Total: total,
		Page:  page,
		Size:  size,
	}, nil
}
