package service

import (
	"context"
	"putra4648/erp/internal/supplier/domain"
	"putra4648/erp/internal/supplier/dto"
	"putra4648/erp/internal/supplier/mapper"
	"putra4648/erp/internal/supplier/repository"

	"github.com/google/uuid"
)

type supplierCommandService struct {
	repo repository.SupplierRepository
}

func NewSupplierCommandService(repo repository.SupplierRepository) SupplierCommandService {
	return &supplierCommandService{repo: repo}
}

func (s *supplierCommandService) Create(ctx context.Context, req *dto.SupplierDTO) (*dto.SupplierDTO, error) {
	supplier := &domain.Supplier{
		ID:      uuid.New(),
		Name:    req.Name,
		Code:    req.Code,
		Address: req.Address,
		Phone:   req.Phone,
		Email:   req.Email,
	}
	if err := s.repo.Save(ctx, supplier); err != nil {
		return nil, err
	}
	return mapper.ToSupplierDTO(supplier), nil
}

func (s *supplierCommandService) Update(ctx context.Context, req *dto.SupplierDTO) (*dto.SupplierDTO, error) {
	supplier, err := s.repo.FindByID(ctx, uuid.Must(uuid.Parse(req.ID)))
	if err != nil {
		return nil, err
	}
	supplier.Name = req.Name
	supplier.Address = req.Address
	supplier.Phone = req.Phone
	supplier.Email = req.Email

	if err := s.repo.Update(ctx, supplier); err != nil {
		return nil, err
	}
	return mapper.ToSupplierDTO(supplier), nil
}

func (s *supplierCommandService) Delete(ctx context.Context, id string) (*dto.SupplierDTO, error) {
	supplierID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	supplier, err := s.repo.FindByID(ctx, supplierID)
	if err != nil {
		return nil, err
	}
	if err := s.repo.Delete(ctx, supplierID); err != nil {
		return nil, err
	}
	return mapper.ToSupplierDTO(supplier), nil
}
