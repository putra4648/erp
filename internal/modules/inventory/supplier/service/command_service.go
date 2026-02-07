package service

import (
	"context"
	"putra4648/erp/internal/modules/inventory/supplier/domain"
	"putra4648/erp/internal/modules/inventory/supplier/dto"
	"putra4648/erp/internal/modules/inventory/supplier/repository"

	"github.com/google/uuid"
)

type SupplierCommandService interface {
	Create(ctx context.Context, req *dto.CreateSupplierRequest) (uuid.UUID, error)
	Update(ctx context.Context, id uuid.UUID, req *dto.UpdateSupplierRequest) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type supplierCommandService struct {
	repo repository.SupplierRepository
}

func NewSupplierCommandService(repo repository.SupplierRepository) SupplierCommandService {
	return &supplierCommandService{repo: repo}
}

func (s *supplierCommandService) Create(ctx context.Context, req *dto.CreateSupplierRequest) (uuid.UUID, error) {
	supplier := &domain.Supplier{
		ID:      uuid.New(),
		Name:    req.Name,
		Code:    req.Code,
		Address: req.Address,
		Phone:   req.Phone,
		Email:   req.Email,
	}
	if err := s.repo.Save(ctx, supplier); err != nil {
		return uuid.Nil, err
	}
	return supplier.ID, nil
}

func (s *supplierCommandService) Update(ctx context.Context, id uuid.UUID, req *dto.UpdateSupplierRequest) error {
	supplier, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	supplier.Name = req.Name
	supplier.Address = req.Address
	supplier.Phone = req.Phone
	supplier.Email = req.Email

	return s.repo.Update(ctx, supplier)
}

func (s *supplierCommandService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
