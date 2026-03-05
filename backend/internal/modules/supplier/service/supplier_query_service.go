package service

import (
	"context"
	"putra4648/erp/internal/modules/supplier/dto"
	"putra4648/erp/internal/modules/supplier/mapper"
	"putra4648/erp/internal/modules/supplier/repository"

	sharedDto "putra4648/erp/internal/modules/shared/dto"

	"github.com/google/uuid"
)

type supplierQueryService struct {
	repo repository.SupplierRepository
}

func NewSupplierQueryService(repo repository.SupplierRepository) SupplierQueryService {
	return &supplierQueryService{repo: repo}
}

func (s *supplierQueryService) FindByID(ctx context.Context, id string) (*dto.SupplierDto, error) {
	supplierID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	supplier, err := s.repo.FindByID(ctx, supplierID)
	if err != nil {
		return nil, err
	}
	return mapper.ToSupplierDto(supplier), nil
}

func (s *supplierQueryService) FindAll(ctx context.Context, req *dto.SupplierFindAllRequest) (*sharedDto.PaginationResponse[*dto.SupplierDto], error) {
	suppliers, total, err := s.repo.FindAll(ctx, req)
	if err != nil {
		return nil, err
	}

	var supplierDtos []*dto.SupplierDto
	for _, supplier := range suppliers {
		supplierDtos = append(supplierDtos, mapper.ToSupplierDto(supplier))
	}

	return &sharedDto.PaginationResponse[*dto.SupplierDto]{
		Items: supplierDtos,
		Total: total,
		Page:  req.Page,
		Size:  req.Size,
	}, nil
}
