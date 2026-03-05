package service

import (
	"context"
	sharedDto "putra4648/erp/internal/modules/shared/dto"
	"putra4648/erp/internal/modules/uom/dto"
	"putra4648/erp/internal/modules/uom/mapper"
	uomRepository "putra4648/erp/internal/modules/uom/repository"

	"github.com/google/uuid"
)

type uomQueryService struct {
	uomRepo uomRepository.UOMRepository
}

func NewUOMQueryService(uomRepo uomRepository.UOMRepository) UOMQueryService {
	return &uomQueryService{uomRepo: uomRepo}
}

func (s *uomQueryService) GetAllUOMs(ctx context.Context, request *dto.UOMRequest) (*sharedDto.PaginationResponse[*dto.UOMDTO], error) {
	uoms, total, err := s.uomRepo.FindAll(ctx, request)
	if err != nil {
		return nil, &UOMError{Code: "DATABASE_ERROR", Message: "Failed to retrieve UOMs"}
	}

	responses := mapper.ToUOMDTOs(uoms)

	return &sharedDto.PaginationResponse[*dto.UOMDTO]{
		Items: responses,
		Total: total,
		Page:  request.Page,
		Size:  request.Size,
	}, nil
}

func (s *uomQueryService) GetUOMByID(ctx context.Context, id uuid.UUID) (*dto.UOMDTO, error) {
	uom, err := s.uomRepo.FindByID(ctx, id)
	if err != nil {
		return nil, &UOMError{Code: "NOT_FOUND", Message: "UOM not found"}
	}

	return mapper.ToUOMDTO(uom), nil
}
