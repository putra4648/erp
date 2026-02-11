package service

import (
	"context"
	"fmt"
	"putra4648/erp/internal/modules/shared/enums"
	"putra4648/erp/internal/modules/stock_movement/domain"
	"putra4648/erp/internal/modules/stock_movement/repository"
	"time"

	"github.com/google/uuid"
)

type StockMovementCommandService interface {
	Create(ctx context.Context, dto *domain.StockMovementDTO) (*domain.StockMovementResponse, error)
	Update(ctx context.Context, id uuid.UUID, dto *domain.StockMovementDTO) (*domain.StockMovementResponse, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type stockMovementCommandService struct {
	repo repository.StockMovementRepository
}

func NewStockMovementCommandService(repo repository.StockMovementRepository) StockMovementCommandService {
	return &stockMovementCommandService{repo: repo}
}

func (s *stockMovementCommandService) Create(ctx context.Context, dto *domain.StockMovementDTO) (*domain.StockMovementResponse, error) {
	dto.ID = uuid.New().String()
	// Generate Movement Number sa-yyyy-mm-dd-xxxx
	dto.MovementNo = fmt.Sprintf("MOV-%s-%d", time.Now().Format("20060102"), time.Now().UnixNano()%10000)

	model := dto.ToModel()
	model.MovementNo = dto.MovementNo // MovementNo is generated server-side usually

	if err := s.repo.Create(ctx, model); err != nil {
		return nil, err
	}

	created, err := s.repo.FindByID(ctx, model.ID)
	if err != nil {
		return nil, err
	}

	return created.ToResponse(), nil
}

func (s *stockMovementCommandService) Update(ctx context.Context, id uuid.UUID, dto *domain.StockMovementDTO) (*domain.StockMovementResponse, error) {
	existing, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if existing.Status != enums.StatusDraft {
		return nil, fmt.Errorf("only draft movements can be updated")
	}

	dto.ID = id.String()
	model := dto.ToModel()
	model.MovementNo = existing.MovementNo

	if err := s.repo.Update(ctx, model); err != nil {
		return nil, err
	}

	updated, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return updated.ToResponse(), nil
}

func (s *stockMovementCommandService) Delete(ctx context.Context, id uuid.UUID) error {
	existing, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if existing.Status != enums.StatusDraft {
		return fmt.Errorf("only draft movements can be deleted")
	}

	return s.repo.Delete(ctx, id)
}
