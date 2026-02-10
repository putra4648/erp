package service

import (
	"context"
	"fmt"
	"putra4648/erp/internal/modules/stock_adjustment/domain"
	"putra4648/erp/internal/modules/stock_adjustment/dto"
	"time"

	"github.com/google/uuid"
)

type StockAdjustmentService interface {
	Create(ctx context.Context, req *dto.CreateStockAdjustmentRequest, userID uuid.UUID) (*dto.StockAdjustmentResponse, error)
	FindByID(ctx context.Context, id uuid.UUID) (*dto.StockAdjustmentResponse, error)
	FindAll(ctx context.Context, page, size int) ([]*dto.StockAdjustmentResponse, int64, error)
}

type stockAdjustmentService struct {
	repo domain.StockAdjustmentRepository
}

func NewStockAdjustmentService(repo domain.StockAdjustmentRepository) StockAdjustmentService {
	return &stockAdjustmentService{repo: repo}
}

func (s *stockAdjustmentService) Create(ctx context.Context, req *dto.CreateStockAdjustmentRequest, userID uuid.UUID) (*dto.StockAdjustmentResponse, error) {
	transactionDate, err := time.Parse("2006-01-02", req.TransactionDate)
	if err != nil {
		return nil, err
	}

	adjustment := &domain.StockAdjustment{
		ID:              uuid.New(),
		AdjustmentNo:    fmt.Sprintf("SA-%d", time.Now().Unix()),
		WarehouseID:     req.WarehouseID,
		TransactionDate: transactionDate,
		Status:          "DRAFT",
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

	return saved.ToResponse(), nil
}

func (s *stockAdjustmentService) FindByID(ctx context.Context, id uuid.UUID) (*dto.StockAdjustmentResponse, error) {
	adjustment, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return adjustment.ToResponse(), nil
}

func (s *stockAdjustmentService) FindAll(ctx context.Context, page, size int) ([]*dto.StockAdjustmentResponse, int64, error) {
	adjustments, total, err := s.repo.FindAll(ctx, page, size)
	if err != nil {
		return nil, 0, err
	}

	var responses []*dto.StockAdjustmentResponse
	for _, adj := range adjustments {
		responses = append(responses, adj.ToResponse())
	}
	return responses, total, nil
}
