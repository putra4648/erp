package mapper

import (
	"putra4648/erp/internal/modules/shared/enums"
	"putra4648/erp/internal/modules/stock_movement/domain"
	"putra4648/erp/internal/modules/stock_movement/dto"
	"time"

	"github.com/google/uuid"
)

func ToModel(dto *dto.StockMovementDTO) *domain.StockMovement {
	transactionDate, _ := time.Parse("2006-01-02", dto.TransactionDate)
	originWarehouseID := uuid.MustParse(dto.OriginWarehouseID)
	destinationWarehouseID := uuid.MustParse(dto.DestinationWarehouseID)

	return &domain.StockMovement{
		ID:                     uuid.MustParse(dto.ID),
		MovementNo:             dto.MovementNo,
		Type:                   enums.TransferType(dto.Type),
		OriginWarehouseID:      &originWarehouseID,
		DestinationWarehouseID: &destinationWarehouseID,
		ReferenceNo:            dto.ReferenceNo,
		Status:                 dto.Status,
		TransactionDate:        transactionDate,
		Note:                   dto.Note,
		Items:                  ToModelItems(dto.Items),
	}
}

func ToModelItems(items []dto.StockMovementItemDTO) []domain.StockMovementItem {
	var result []domain.StockMovementItem
	for _, item := range items {
		result = append(result, *ToModelItem(&item))
	}
	return result
}

func ToModelItem(dto *dto.StockMovementItemDTO) *domain.StockMovementItem {
	return &domain.StockMovementItem{
		ID:        uuid.MustParse(dto.ID),
		ProductID: uuid.MustParse(dto.ProductID),
		Quantity:  dto.Quantity,
		Note:      dto.Note,
	}
}

func ToDTO(model *domain.StockMovement) *dto.StockMovementDTO {
	return &dto.StockMovementDTO{
		ID:                     model.ID.String(),
		MovementNo:             model.MovementNo,
		Type:                   string(model.Type),
		OriginWarehouseID:      model.OriginWarehouseID.String(),
		DestinationWarehouseID: model.DestinationWarehouseID.String(),
		ReferenceNo:            model.ReferenceNo,
		Status:                 model.Status,
		TransactionDate:        model.TransactionDate.Format("2006-01-02"),
		Note:                   model.Note,
		Items:                  ToDTOItems(model.Items),
	}
}

func ToDTOItems(items []domain.StockMovementItem) []dto.StockMovementItemDTO {
	var result []dto.StockMovementItemDTO
	for _, item := range items {
		result = append(result, *ToDTOItem(&item))
	}
	return result
}

func ToDTOItem(model *domain.StockMovementItem) *dto.StockMovementItemDTO {
	return &dto.StockMovementItemDTO{
		ID:        model.ID.String(),
		ProductID: model.ProductID.String(),
		Quantity:  model.Quantity,
		Note:      model.Note,
	}
}
