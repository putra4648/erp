package mapper

import (
	"putra4648/erp/internal/supplier/domain"
	"putra4648/erp/internal/supplier/dto"
)

func ToSupplierDto(supplier *domain.Supplier) *dto.SupplierDto {
	return &dto.SupplierDto{
		ID:      supplier.ID.String(),
		Name:    supplier.Name,
		Code:    supplier.Code,
		Address: supplier.Address,
		Phone:   supplier.Phone,
		Email:   supplier.Email,
	}
}

func ToSupplier(req *dto.SupplierDto) *domain.Supplier {
	return &domain.Supplier{
		Name:    req.Name,
		Code:    req.Code,
		Address: req.Address,
		Phone:   req.Phone,
		Email:   req.Email,
	}
}
