package service

import (
	"context"
	"fmt"
	"putra4648/erp/internal/shared/enums"
	stockLevelService "putra4648/erp/internal/stock_level/service"
	"putra4648/erp/internal/stock_movement/domain"
	"putra4648/erp/internal/stock_movement/dto"
	"putra4648/erp/internal/stock_movement/mapper"
	"putra4648/erp/internal/stock_movement/repository"
	"time"

	"github.com/google/uuid"
)

type stockMovementCommandService struct {
	repo              repository.StockMovementRepository
	stockLevelService stockLevelService.StockLevelCommandService
}

func NewStockMovementCommandService(repo repository.StockMovementRepository, stockLevelService stockLevelService.StockLevelCommandService) StockMovementCommandService {
	return &stockMovementCommandService{
		repo:              repo,
		stockLevelService: stockLevelService,
	}
}

func (s *stockMovementCommandService) Create(ctx context.Context, dto *dto.StockMovementDTO) (*dto.StockMovementDTO, error) {
	dto.ID = uuid.New().String()
	// Generate Movement Number sa-yyyy-mm-dd-xxxx
	dto.MovementNo = fmt.Sprintf("MOV-%s-%d", time.Now().Format("20060102"), time.Now().UnixNano()%10000)

	model := mapper.ToModel(dto)
	model.MovementNo = dto.MovementNo // MovementNo is generated server-side usually

	if err := s.repo.Create(ctx, model); err != nil {
		return nil, err
	}

	created, err := s.repo.FindByID(ctx, model.ID)
	if err != nil {
		return nil, err
	}

	return mapper.ToDTO(created), nil
}

func (s *stockMovementCommandService) Update(ctx context.Context, id uuid.UUID, dto *dto.StockMovementDTO) (*dto.StockMovementDTO, error) {
	existing, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if existing.Status != enums.StatusDraft {
		return nil, fmt.Errorf("only draft movements can be updated")
	}

	dto.ID = id.String()
	model := mapper.ToModel(dto)
	model.MovementNo = existing.MovementNo

	if err := s.repo.Update(ctx, model); err != nil {
		return nil, err
	}

	updated, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapper.ToDTO(updated), nil
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

func (s *stockMovementCommandService) Approve(ctx context.Context, id uuid.UUID) error {
	movement, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	// ... validasi status movement ...

	for _, item := range movement.Items {
		// PROSES OUTGOING (Gudang Asal)
		if movement.OriginWarehouseID != nil {
			// 1. Update Stock Level (Atomic UPDATE quantity = quantity + delta)
			err := s.stockLevelService.AdjustStock(ctx, item.ProductID, *movement.OriginWarehouseID, item.Quantity.Neg())
			if err != nil {
				return err
			}

			// 2. Insert Transaction (Memicu Trigger AFTER INSERT)
			transaction := &domain.StockTransaction{
				ID:          uuid.New(),
				ProductID:   item.ProductID,
				WarehouseID: *movement.OriginWarehouseID,
				SupplierID:  item.Product.SupplierID,
				Type:        "OUT",
				Quantity:    item.Quantity.Neg(),
				ReferenceNo: movement.MovementNo,
			}
			if err := s.repo.CreateTransaction(ctx, transaction); err != nil {
				return err
			}
		}

		// PROSES INCOMING (Gudang Tujuan)
		if movement.DestinationWarehouseID != nil {
			// 1. Update Stock Level
			err := s.stockLevelService.AdjustStock(ctx, item.ProductID, *movement.DestinationWarehouseID, item.Quantity)
			if err != nil {
				return err
			}

			// 2. Insert Transaction (Memicu Trigger AFTER INSERT)
			transaction := &domain.StockTransaction{
				ID:          uuid.New(),
				ProductID:   item.ProductID,
				WarehouseID: *movement.DestinationWarehouseID,
				SupplierID:  item.Product.SupplierID,
				Type:        "IN",
				Quantity:    item.Quantity,
				ReferenceNo: movement.MovementNo,
			}
			if err := s.repo.CreateTransaction(ctx, transaction); err != nil {
				return err
			}
		}
	}

	return s.repo.CompletedMovement(ctx, id)
}
