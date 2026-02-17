package service

import (
	"context"
	uomDto "putra4648/erp/internal/modules/uom/dto"
	"putra4648/erp/internal/modules/uom/mapper"
	uomRepository "putra4648/erp/internal/modules/uom/repository"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

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

func (s *uomCommandService) CreateUOM(ctx context.Context, uomDTO *uomDto.UOMDTO) (*uomDto.UOMDTO, error) {
	uom := mapper.ToUOM(uomDTO)

	err := s.uomRepo.Create(ctx, uom)
	if err != nil {
		s.logger.Error("Failed to create UOM in DB", zap.Error(err), zap.String("name", uom.Name))
		return nil, &UOMError{Code: "DATABASE_ERROR", Message: "Failed to create UOM"}
	}

	return mapper.ToUOMDTO(uom), nil
}

func (s *uomCommandService) UpdateUOM(ctx context.Context, id uuid.UUID, uomDTO *uomDto.UOMDTO) (*uomDto.UOMDTO, error) {
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

	return mapper.ToUOMDTO(existingUOM), nil
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
