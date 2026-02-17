package mapper

import (
	"putra4648/erp/internal/modules/uom/domain"
	"putra4648/erp/internal/modules/uom/dto"
)

func ToUOM(uomDTO *dto.UOMDTO) *domain.UOM {
	return &domain.UOM{
		Name: uomDTO.Name,
	}
}

func ToUOMs(uoms []*dto.UOMDTO) []*domain.UOM {
	responses := make([]*domain.UOM, len(uoms))
	for i, uom := range uoms {
		responses[i] = ToUOM(uom)
	}
	return responses
}

func ToUOMDTO(uom *domain.UOM) *dto.UOMDTO {
	return &dto.UOMDTO{
		ID:   uom.ID.String(),
		Name: uom.Name,
	}
}

func ToUOMDTOs(uoms []*domain.UOM) []*dto.UOMDTO {
	responses := make([]*dto.UOMDTO, len(uoms))
	for i, uom := range uoms {
		responses[i] = ToUOMDTO(uom)
	}
	return responses
}
