package service

import (
	uomModel "putra4648/erp/internal/modules/uom/model"
	uomRepository "putra4648/erp/internal/modules/uom/repository"

	"github.com/google/uuid"
)

type UOMQueryService interface {
	GetAllUOMs() ([]*uomModel.UOMResponse, error)
	GetUOMByID(id uuid.UUID) (*uomModel.UOMResponse, error)
}

type uomQueryService struct {
	uomRepo uomRepository.UOMRepository
}

func NewUOMQueryService(uomRepo uomRepository.UOMRepository) UOMQueryService {
	return &uomQueryService{uomRepo: uomRepo}
}

func (s *uomQueryService) GetAllUOMs() ([]*uomModel.UOMResponse, error) {
	uoms, err := s.uomRepo.FindAll()
	if err != nil {
		return nil, &UOMError{Code: "DATABASE_ERROR", Message: "Failed to retrieve UOMs"}
	}

	responses := make([]*uomModel.UOMResponse, len(uoms))
	for i, uom := range uoms {
		responses[i] = uom.ToResponse()
	}

	return responses, nil
}

func (s *uomQueryService) GetUOMByID(id uuid.UUID) (*uomModel.UOMResponse, error) {
	uom, err := s.uomRepo.FindByID(id)
	if err != nil {
		return nil, &UOMError{Code: "NOT_FOUND", Message: "UOM not found"}
	}

	return uom.ToResponse(), nil
}
