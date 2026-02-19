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

	var originWarehouseID *uuid.UUID
	if dto.OriginWarehouseID != "" {
		id, err := uuid.Parse(dto.OriginWarehouseID)
		if err == nil {
			originWarehouseID = &id
		}
	}

	var destinationWarehouseID *uuid.UUID
	if dto.DestinationWarehouseID != "" {
		id, err := uuid.Parse(dto.DestinationWarehouseID)
		if err == nil {
			destinationWarehouseID = &id
		}
	}

	return &domain.StockMovement{
		ID:                     uuid.MustParse(dto.ID),
		MovementNo:             dto.MovementNo,
		Type:                   enums.TransferType(dto.Type),
		OriginWarehouseID:      originWarehouseID,
		DestinationWarehouseID: destinationWarehouseID,
		ReferenceNo:            dto.ReferenceNo,
		Status:                 dto.Status,
		TransactionDate:        transactionDate,
		Note:                   dto.Note,
		Items:                  ToModelItems(dto.Items),
	}
}

func ToModelItems(items []*dto.StockMovementItemDTO) []*domain.StockMovementItem {
	var result []*domain.StockMovementItem
	for _, item := range items {
		result = append(result, ToModelItem(item))
	}
	return result
}

func ToModelItem(dto *dto.StockMovementItemDTO) *domain.StockMovementItem {
	id := uuid.Nil
	if dto.ID != "" {
		id, _ = uuid.Parse(dto.ID)
	}
	if id == uuid.Nil {
		id = uuid.New()
	}

	return &domain.StockMovementItem{
		ID:        id,
		ProductID: uuid.MustParse(dto.ProductID),
		Quantity:  dto.Quantity,
		Note:      dto.Note,
	}
}

func ToDTO(model *domain.StockMovement) *dto.StockMovementDTO {
	originWarehouseID := ""
	if model.OriginWarehouseID != nil {
		originWarehouseID = model.OriginWarehouseID.String()
	}

	destinationWarehouseID := ""
	if model.DestinationWarehouseID != nil {
		destinationWarehouseID = model.DestinationWarehouseID.String()
	}

	return &dto.StockMovementDTO{
		ID:                     model.ID.String(),
		MovementNo:             model.MovementNo,
		Type:                   string(model.Type),
		OriginWarehouseID:      originWarehouseID,
		DestinationWarehouseID: destinationWarehouseID,
		ReferenceNo:            model.ReferenceNo,
		Status:                 model.Status,
		TransactionDate:        model.TransactionDate.Format("2006-01-02"),
		Note:                   model.Note,
		Items:                  ToDTOItems(model.Items),
	}
}

func ToDTOItems(items []*domain.StockMovementItem) []*dto.StockMovementItemDTO {
	result := make([]*dto.StockMovementItemDTO, 0)
	for _, item := range items {
		if item != nil {
			result = append(result, ToDTOItem(item))
		}
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

func ToTransactionDTO(model *domain.StockTransaction) *dto.StockTransactionResponse {

	return &dto.StockTransactionResponse{
		ID:            model.ID.String(),
		ProductID:     model.ProductID.String(),
		ProductName:   model.Product.Name,
		WarehouseID:   model.WarehouseID.String(),
		WarehouseName: model.Warehouse.Name,
		SupplierID:    model.SupplierID.String(),
		SupplierName:  model.Supplier.Name,
		Type:          model.Type,
		Quantity:      model.Quantity,
		ReferenceNo:   model.ReferenceNo,
		CreatedAt:     model.CreatedAt.Format(time.RFC3339),
	}
}

func ToTransactionDTOs(models []*domain.StockTransaction) []*dto.StockTransactionResponse {
	result := make([]*dto.StockTransactionResponse, len(models))
	for i, m := range models {
		result[i] = ToTransactionDTO(m)
	}
	return result
}
