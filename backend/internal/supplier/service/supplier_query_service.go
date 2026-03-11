package service

import (
	"context"
	"putra4648/erp/internal/supplier/dto"
	"putra4648/erp/internal/supplier/mapper"
	"putra4648/erp/internal/supplier/repository"

	sharedDto "putra4648/erp/internal/shared/dto"

	"github.com/google/uuid"
)

type supplierQueryService struct {
	repo repository.SupplierRepository
}

func NewSupplierQueryService(repo repository.SupplierRepository) SupplierQueryService {
	return &supplierQueryService{repo: repo}
}

func (s *supplierQueryService) FindByID(ctx context.Context, id string) (*dto.SupplierDTO, error) {
	supplierID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	supplier, err := s.repo.FindByID(ctx, supplierID)
	if err != nil {
		return nil, err
	}
	return mapper.ToSupplierDTO(supplier), nil
}

func (s *supplierQueryService) FindAll(ctx context.Context, pagination *sharedDto.PaginationRequest, req *dto.SupplierDTO) (*sharedDto.PaginationResponse[*dto.SupplierDTO], error) {
	suppliers, total, err := s.repo.FindAll(ctx, pagination, req)
	if err != nil {
		return nil, err
	}

	var supplierDtos []*dto.SupplierDTO
	for _, supplier := range suppliers {
		supplierDtos = append(supplierDtos, mapper.ToSupplierDTO(supplier))
	}

	return &sharedDto.PaginationResponse[*dto.SupplierDTO]{
		Items: supplierDtos,
		Total: total,
		Page:  pagination.Page,
		Size:  pagination.Size,
	}, nil
}
