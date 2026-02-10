package service

import (
	"putra4648/erp/configs/logger"
	"putra4648/erp/internal/modules/inventory/stock/domain"
	"putra4648/erp/internal/modules/inventory/stock/dto"
	approvalDomain "putra4648/erp/internal/modules/shared/approval/domain"

	"github.com/google/uuid"
	"go.uber.org/dig"
)

type ApprovalRepository interface {
	Create(docCode string, referenceID uuid.UUID) (*approvalDomain.ApprovalTransaction, error)
}

type StockRepository interface {
	CreateStockAdjustment(dto dto.StockRequest, userID uuid.UUID) (*domain.StockAdjustment, error)
}
type StockService struct {
	repository         StockRepository
	approvalRepository ApprovalRepository
}

type StockServiceParams struct {
	dig.In

	Repository         StockRepository
	ApprovalRepository ApprovalRepository
}

func NewStockCommandService(params StockServiceParams) *StockService {
	return &StockService{
		repository:         params.Repository,
		approvalRepository: params.ApprovalRepository,
	}
}

func (s *StockService) CreateStockAdjustment(dto dto.StockRequest, userId uuid.UUID) error {
	// first create stock adjustment
	res, err := s.repository.CreateStockAdjustment(dto, userId)
	if err != nil {
		return err
	}

	// then create new approval
	approvalRes, approvalErr := s.approvalRepository.Create("STOCK_ADJUSTMENT", res.ID)
	if approvalErr != nil {
		return approvalErr
	}

	logger.Log.Info("Approval result %v ", approvalRes)

	return approvalErr

}
