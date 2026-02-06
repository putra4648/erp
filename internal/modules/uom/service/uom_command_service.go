package service

import (
	uomModel "putra4648/erp/internal/modules/uom/model"
	uomRepository "putra4648/erp/internal/modules/uom/repository"

	"github.com/google/uuid"
)

type UOMCommandService interface {
	CreateUOM(uomDTO *uomModel.UOMDTO) (*uomModel.UOMResponse, error)
	UpdateUOM(id uuid.UUID, uomDTO *uomModel.UOMDTO) (*uomModel.UOMResponse, error)
	DeleteUOM(id uuid.UUID) error
}

type uomCommandService struct {
	uomRepo uomRepository.UOMRepository
}

func NewUOMCommandService(uomRepo uomRepository.UOMRepository) UOMCommandService {
	return &uomCommandService{uomRepo: uomRepo}
}

func (s *uomCommandService) CreateUOM(uomDTO *uomModel.UOMDTO) (*uomModel.UOMResponse, error) {
	uom := uomDTO.ToModel()

	err := s.uomRepo.Create(uom)
	if err != nil {
		return nil, &UOMError{Code: "DATABASE_ERROR", Message: "Failed to create UOM"}
	}

	return uom.ToResponse(), nil
}

func (s *uomCommandService) UpdateUOM(id uuid.UUID, uomDTO *uomModel.UOMDTO) (*uomModel.UOMResponse, error) {
	existingUOM, err := s.uomRepo.FindByID(id)
	if err != nil {
		return nil, &UOMError{Code: "NOT_FOUND", Message: "UOM not found"}
	}

	existingUOM.Name = uomDTO.Name

	err = s.uomRepo.Update(existingUOM)
	if err != nil {
		return nil, &UOMError{Code: "DATABASE_ERROR", Message: "Failed to update UOM"}
	}

	return existingUOM.ToResponse(), nil
}

func (s *uomCommandService) DeleteUOM(id uuid.UUID) error {
	_, err := s.uomRepo.FindByID(id)
	if err != nil {
		return &UOMError{Code: "NOT_FOUND", Message: "UOM not found"}
	}

	err = s.uomRepo.Delete(id)
	if err != nil {
		return &UOMError{Code: "DATABASE_ERROR", Message: "Failed to delete UOM"}
	}

	return nil
}
