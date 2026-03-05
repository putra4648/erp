package repository

import (
	"context"
	"putra4648/erp/internal/modules/stock_movement/domain"
	"putra4648/erp/internal/modules/stock_movement/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type stockMovementRepository struct {
	db *gorm.DB
}

func NewStockMovementRepository(db *gorm.DB) StockMovementRepository {
	return &stockMovementRepository{db: db}
}

func (r *stockMovementRepository) Create(ctx context.Context, movement *domain.StockMovement) error {
	return r.db.WithContext(ctx).Create(movement).Error
}

func (r *stockMovementRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.StockMovement, error) {
	var movement domain.StockMovement
	err := r.db.WithContext(ctx).
		Preload("OriginWarehouse").
		Preload("DestinationWarehouse").
		Preload("Items.Product").
		Where("id = ?", id).First(&movement).Error
	if err != nil {
		return nil, err
	}
	return &movement, nil
}

func (r *stockMovementRepository) FindAll(ctx context.Context, req *dto.StockMovementRequest) ([]*domain.StockMovement, int64, error) {
	var movements []*domain.StockMovement
	var total int64
	db := r.db.WithContext(ctx).Model(&domain.StockMovement{}).
		Preload("OriginWarehouse").
		Preload("DestinationWarehouse").
		Preload("Items")

	if req.Type != "" {
		db = db.Where("type = ?", req.Type)
	}

	if req.Search != "" {
		searchTerm := "%" + req.Search + "%"
		db = db.Where("movement_no LIKE ? OR reference_no LIKE ? OR note LIKE ?", searchTerm, searchTerm, searchTerm)
	}

	db.Count(&total)

	if req.Page > 0 && req.Size > 0 {
		offset := (req.Page - 1) * req.Size
		db = db.Limit(req.Size).Offset(offset)
	}

	err := db.Order("transaction_date desc").Find(&movements).Error
	if err != nil {
		return nil, 0, err
	}
	return movements, total, nil
}

func (r *stockMovementRepository) Update(ctx context.Context, movement *domain.StockMovement) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(movement).Association("Items").Replace(movement.Items); err != nil {
			return err
		}
		return tx.Save(movement).Error
	})
}

func (r *stockMovementRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&domain.StockMovement{}, "id = ?", id).Error
}

func (r *stockMovementRepository) CompletedMovement(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Model(&domain.StockMovement{}).Where("id = ?", id).Update("status", "COMPLETED").Error
}

func (r *stockMovementRepository) CreateTransaction(ctx context.Context, transaction *domain.StockTransaction) error {
	return r.db.WithContext(ctx).Create(transaction).Error
}

func (r *stockMovementRepository) FindTransactions(ctx context.Context, req *dto.StockTransactionRequest) ([]*domain.StockTransaction, int64, error) {
	var transactions []*domain.StockTransaction
	var total int64
	query := r.db.WithContext(ctx).Model(&domain.StockTransaction{}).
		Preload("Product").
		Preload("Warehouse").
		Preload("Supplier").
		Order("created_at DESC")

	if req.ProductID != "" {
		query = query.Where("product_id = ?", req.ProductID)
	}
	if req.WarehouseID != "" {
		query = query.Where("warehouse_id = ?", req.WarehouseID)
	}

	query.Count(&total)

	if req.Page > 0 && req.Size > 0 {
		offset := (req.Page - 1) * req.Size
		query = query.Limit(req.Size).Offset(offset)
	}

	err := query.Find(&transactions).Error
	return transactions, total, err
}
