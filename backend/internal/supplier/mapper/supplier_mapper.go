package mapper

import (
	"putra4648/erp/internal/supplier/domain"
	"putra4648/erp/internal/supplier/dto"
)

func ToSupplierDTO(supplier *domain.Supplier) *dto.SupplierDTO {
	return &dto.SupplierDTO{
		ID:      supplier.ID.String(),
		Name:    supplier.Name,
		Code:    supplier.Code,
		Address: supplier.Address,
		Phone:   supplier.Phone,
		Email:   supplier.Email,
	}
}

func ToSupplier(req *dto.SupplierDTO) *domain.Supplier {
	return &domain.Supplier{
		Name:    req.Name,
		Code:    req.Code,
		Address: req.Address,
		Phone:   req.Phone,
		Email:   req.Email,
	}
}
