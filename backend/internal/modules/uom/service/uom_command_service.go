package service

import (
	"context"
	uomDomain "putra4648/erp/internal/modules/uom/domain"
	uomRepository "putra4648/erp/internal/modules/uom/repository"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type UOMCommandService interface {
	CreateUOM(ctx context.Context, uomDTO *uomDomain.UOMDTO) (*uomDomain.UOMResponse, error)
	UpdateUOM(ctx context.Context, id uuid.UUID, uomDTO *uomDomain.UOMDTO) (*uomDomain.UOMResponse, error)
	DeleteUOM(ctx context.Context, id uuid.UUID) error
}

type uomCommandService struct {
	uomRepo uomRepository.UOMRepository
	logger  *zap.Logger
}

func NewUOMCommandService(uomRepo uomRepository.UOMRepository, logger *zap.Logger) UOMCommandService {
	return &uomCommandService{
		uomRepo: uomRepo,
		logger:  logger,
	}
}

func (s *uomCommandService) CreateUOM(ctx context.Context, uomDTO *uomDomain.UOMDTO) (*uomDomain.UOMResponse, error) {
	uom := uomDTO.ToModel()

	err := s.uomRepo.Create(ctx, uom)
	if err != nil {
		s.logger.Error("Failed to create UOM in DB", zap.Error(err), zap.String("name", uom.Name))
		return nil, &UOMError{Code: "DATABASE_ERROR", Message: "Failed to create UOM"}
	}

	return uom.ToResponse(), nil
}

func (s *uomCommandService) UpdateUOM(ctx context.Context, id uuid.UUID, uomDTO *uomDomain.UOMDTO) (*uomDomain.UOMResponse, error) {
	existingUOM, err := s.uomRepo.FindByID(ctx, id)
	if err != nil {
		return nil, &UOMError{Code: "NOT_FOUND", Message: "UOM not found"}
	}

	existingUOM.Name = uomDTO.Name

	err = s.uomRepo.Update(ctx, existingUOM)
	if err != nil {
		s.logger.Error("Failed to update UOM in DB", zap.Error(err), zap.String("id", id.String()))
		return nil, &UOMError{Code: "DATABASE_ERROR", Message: "Failed to update UOM"}
	}

	return existingUOM.ToResponse(), nil
}

func (s *uomCommandService) DeleteUOM(ctx context.Context, id uuid.UUID) error {
	_, err := s.uomRepo.FindByID(ctx, id)
	if err != nil {
		return &UOMError{Code: "NOT_FOUND", Message: "UOM not found"}
	}

	err = s.uomRepo.Delete(ctx, id)
	if err != nil {
		s.logger.Error("Failed to delete UOM in DB", zap.Error(err), zap.String("id", id.String()))
		return &UOMError{Code: "DATABASE_ERROR", Message: "Failed to delete UOM"}
	}

	return nil
}
