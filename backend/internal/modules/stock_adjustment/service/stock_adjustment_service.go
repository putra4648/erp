package service

import (
	"context"
	"fmt"
	"putra4648/erp/internal/modules/shared/enums"
	"putra4648/erp/internal/modules/stock_adjustment/domain"
	"putra4648/erp/internal/modules/stock_adjustment/dto"
	"putra4648/erp/internal/modules/stock_adjustment/mapper"
	"putra4648/erp/internal/modules/stock_adjustment/repository"
	"time"

	"github.com/google/uuid"
)

type stockAdjustmentService struct {
	repo repository.StockAdjustmentRepository
}

func NewStockAdjustmentService(repo repository.StockAdjustmentRepository) StockAdjustmentService {
	return &stockAdjustmentService{repo: repo}
}

func (s *stockAdjustmentService) Create(ctx context.Context, req *dto.CreateStockAdjustmentRequest, userID uuid.UUID) (*dto.StockAdjustmentDto, error) {
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

func (s *stockAdjustmentService) FindByID(ctx context.Context, id uuid.UUID) (*dto.StockAdjustmentDto, error) {
	adjustment, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapper.ToStockAdjustmentDto(adjustment), nil
}

func (s *stockAdjustmentService) FindAll(ctx context.Context, page, size int) ([]*dto.StockAdjustmentDto, int64, error) {
	adjustments, total, err := s.repo.FindAll(ctx, page, size)
	if err != nil {
		return nil, 0, err
	}

	var responses []*dto.StockAdjustmentDto
	for _, adj := range adjustments {
		responses = append(responses, mapper.ToStockAdjustmentDto(adj))
	}
	return responses, total, nil
}

func (s *stockAdjustmentService) Update(ctx context.Context, id uuid.UUID, req *dto.CreateStockAdjustmentRequest) (*dto.StockAdjustmentDto, error) {
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

func (s *stockAdjustmentService) Approve(ctx context.Context, id uuid.UUID, userID uuid.UUID) (*dto.StockAdjustmentDto, error) {
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

func (s *stockAdjustmentService) Void(ctx context.Context, id uuid.UUID) (*dto.StockAdjustmentDto, error) {
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
