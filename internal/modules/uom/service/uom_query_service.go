package service

import (
	"context"
	uomDomain "putra4648/erp/internal/modules/uom/domain"
	"putra4648/erp/internal/modules/uom/dto"
	uomRepository "putra4648/erp/internal/modules/uom/repository"
	sharedDto "putra4648/erp/internal/modules/shared/dto"

	"github.com/google/uuid"
)

type UOMQueryService interface {
	GetAllUOMs(ctx context.Context, request *dto.UOMRequest) (*sharedDto.PaginationResponse[*uomDomain.UOMResponse], error)
	GetUOMByID(ctx context.Context, id uuid.UUID) (*uomDomain.UOMResponse, error)
}

type uomQueryService struct {
	uomRepo uomRepository.UOMRepository
}

func NewUOMQueryService(uomRepo uomRepository.UOMRepository) UOMQueryService {
	return &uomQueryService{uomRepo: uomRepo}
}

func (s *uomQueryService) GetAllUOMs(ctx context.Context, request *dto.UOMRequest) (*sharedDto.PaginationResponse[*uomDomain.UOMResponse], error) {
	uoms, total, err := s.uomRepo.FindAll(ctx, request)
	if err != nil {
		return nil, &UOMError{Code: "DATABASE_ERROR", Message: "Failed to retrieve UOMs"}
	}

	responses := make([]*uomDomain.UOMResponse, len(uoms))
	for i, uom := range uoms {
		responses[i] = uom.ToResponse()
	}

	return &sharedDto.PaginationResponse[*uomDomain.UOMResponse]{
		Items: responses,
		Total: total,
		Page:  request.Page,
		Size:  request.Size,
	}, nil
}

func (s *uomQueryService) GetUOMByID(ctx context.Context, id uuid.UUID) (*uomDomain.UOMResponse, error) {
	uom, err := s.uomRepo.FindByID(ctx, id)
	if err != nil {
		return nil, &UOMError{Code: "NOT_FOUND", Message: "UOM not found"}
	}

	return uom.ToResponse(), nil
}
