package service

import (
	"putra4648/erp/internal/modules/product/model"
	"putra4648/erp/internal/modules/product/repository"
)

type UOMQueryService interface {
	GetAllUOMs() ([]*model.UOMResponse, error)
}

type uomQueryService struct {
	uomRepo repository.UOMRepository
}

func NewUOMQueryService(uomRepo repository.UOMRepository) UOMQueryService {
	return &uomQueryService{uomRepo: uomRepo}
}

func (s *uomQueryService) GetAllUOMs() ([]*model.UOMResponse, error) {
	uoms, err := s.uomRepo.FindAll()
	if err != nil {
		return nil, &UOMError{Code: "DATABASE_ERROR", Message: "Failed to retrieve UOMs"}
	}

	responses := make([]*model.UOMResponse, len(uoms))
	for i, uom := range uoms {
		responses[i] = uom.ToResponse()
	}

	return responses, nil
}
