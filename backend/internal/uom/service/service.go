package service

import (
	"context"
	sharedDto "putra4648/erp/internal/shared/dto"
	"putra4648/erp/internal/uom/dto"
	uomDto "putra4648/erp/internal/uom/dto"

	"github.com/google/uuid"
)

type UOMCommandService interface {
	CreateUOM(ctx context.Context, uomDTO *dto.UOMDTO) (*dto.UOMDTO, error)
	UpdateUOM(ctx context.Context, id uuid.UUID, uomDTO *uomDto.UOMDTO) (*uomDto.UOMDTO, error)
	DeleteUOM(ctx context.Context, id uuid.UUID) error
}

type UOMQueryService interface {
	GetAllUOMs(ctx context.Context, request *dto.UOMRequest) (*sharedDto.PaginationResponse[*uomDto.UOMDTO], error)
	GetUOMByID(ctx context.Context, id uuid.UUID) (*uomDto.UOMDTO, error)
}
