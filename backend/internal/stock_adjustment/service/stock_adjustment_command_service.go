package service

import (
	"context"
	"fmt"
	"putra4648/erp/internal/shared/enums"
	"putra4648/erp/internal/stock_adjustment/domain"
	"putra4648/erp/internal/stock_adjustment/dto"
	"putra4648/erp/internal/stock_adjustment/mapper"
	"putra4648/erp/internal/stock_adjustment/repository"
	"time"

	"github.com/google/uuid"
)

type stockAdjustmentCommandService struct {
	repo repository.StockAdjustmentRepository
}

func NewStockAdjustmentCommandService(repo repository.StockAdjustmentRepository) StockAdjustmentCommandService {
	return &stockAdjustmentCommandService{repo: repo}
}

func (s *stockAdjustmentCommandService) Create(ctx context.Context, req *dto.CreateStockAdjustmentRequest, userID uuid.UUID) (*dto.StockAdjustmentDto, error) {
	transactionDate, err := time.Parse("2006-01-02", req.TransactionDate)
	if err != nil {
		return nil, err
	}

	adjustment := &domain.StockAdjustment{
		ID:              uuid.New(),
		AdjustmentNo:    fmt.Sprintf("SA-%d", time.Now().Unix()),
		WarehouseID:     req.WarehouseID,
		TransactionDate: transactionDate,
		Status:          enums.StatusDraft,
		Note:            req.Note,
		CreatedBy:       userID,
	}

	for _, itemReq := range req.Items {
		adjustment.Items = append(adjustment.Items, domain.StockAdjustmentItem{
			ID:            uuid.New(),
			ProductID:     itemReq.ProductID,
			ReasonID:      itemReq.ReasonID,
			ActualQty:     itemReq.ActualQty,
			SystemQty:     itemReq.SystemQty,
			AdjustmentQty: itemReq.ActualQty.Sub(itemReq.SystemQty),
		})
	}

	if err := s.repo.Save(ctx, adjustment); err != nil {
		return nil, err
	}

	// Reload to get preloaded data
	saved, err := s.repo.FindByID(ctx, adjustment.ID)
	if err != nil {
		return nil, err
	}

	return mapper.ToStockAdjustmentDto(saved), nil
}

func (s *stockAdjustmentCommandService) Update(ctx context.Context, id uuid.UUID, req *dto.CreateStockAdjustmentRequest) (*dto.StockAdjustmentDto, error) {
	adjustment, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if adjustment.Status != enums.StatusDraft {
		return nil, fmt.Errorf("only draft adjustment can be updated")
	}

	transactionDate, err := time.Parse("2006-01-02", req.TransactionDate)
	if err != nil {
		return nil, err
	}

	adjustment.WarehouseID = req.WarehouseID
	adjustment.TransactionDate = transactionDate
	adjustment.Note = req.Note

	// Replace items
	var items []domain.StockAdjustmentItem
	for _, itemReq := range req.Items {
		items = append(items, domain.StockAdjustmentItem{
			ID:                uuid.New(),
			StockAdjustmentID: id,
			ProductID:         itemReq.ProductID,
			ReasonID:          itemReq.ReasonID,
			ActualQty:         itemReq.ActualQty,
			SystemQty:         itemReq.SystemQty,
			AdjustmentQty:     itemReq.ActualQty.Sub(itemReq.SystemQty),
		})
	}
	adjustment.Items = items

	if err := s.repo.Update(ctx, adjustment); err != nil {
		return nil, err
	}

	saved, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapper.ToStockAdjustmentDto(saved), nil
}

func (s *stockAdjustmentCommandService) Approve(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*dto.StockAdjustmentDto, error) {
	adjustment, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if adjustment.Status != enums.StatusDraft {
		return nil, fmt.Errorf("only draft adjustment can be approved")
	}

	adjustment.Status = enums.StatusApproved
	adjustment.ApprovedBy = &userID

	if err := s.repo.Update(ctx, adjustment); err != nil {
		return nil, err
	}

	saved, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapper.ToStockAdjustmentDto(saved), nil
}

func (s *stockAdjustmentCommandService) Void(ctx context.Context, id uuid.UUID) (*dto.StockAdjustmentDto, error) {
	adjustment, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if adjustment.Status != enums.StatusDraft && adjustment.Status != enums.StatusApproved {
		return nil, fmt.Errorf("only draft or approved adjustment can be voided")
	}

	adjustment.Status = enums.StatusVoid

	if err := s.repo.Update(ctx, adjustment); err != nil {
		return nil, err
	}

	saved, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapper.ToStockAdjustmentDto(saved), nil
}
