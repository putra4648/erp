package repository

import (
	"context"
	sharedDto "putra4648/erp/internal/shared/dto"
	"putra4648/erp/internal/stock_movement/domain"
	"putra4648/erp/internal/stock_movement/dto"

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

func (r *stockMovementRepository) FindAll(ctx context.Context, pagination *sharedDto.PaginationRequest, req *dto.StockMovementDTO) ([]*domain.StockMovement, int64, error) {
	var movements []*domain.StockMovement
	var total int64

	db := r.db.WithContext(ctx).Model(&domain.StockMovement{}).
		Preload("OriginWarehouse").
		Preload("DestinationWarehouse").
		Preload("Items")

	if req.Type != "" {
		db = db.Where("type = ?", req.Type)
	}

	if req.Name != "" {
		searchTerm := "%" + req.Name + "%"
		db = db.Where("movement_no LIKE ? OR reference_no LIKE ? OR note LIKE ?", searchTerm, searchTerm, searchTerm)
	}

	db.Count(&total)

	if pagination.Page > 0 && pagination.Size > 0 {
		offset := (pagination.Page - 1) * pagination.Size
		db = db.Limit(pagination.Size).Offset(offset)
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

func (r *stockMovementRepository) FindTransactions(ctx context.Context, pagination *sharedDto.PaginationRequest, req *dto.StockTransactionDTO) ([]*domain.StockTransaction, int64, error) {
	var transactions []*domain.StockTransaction
	var total int64
	db := r.db.WithContext(ctx).Model(&domain.StockTransaction{}).
		Preload("Product").
		Preload("Warehouse").
		Preload("Supplier").
		Order("created_at DESC")

	if req.ProductID != "" {
		db = db.Where("product_id = ?", req.ProductID)
	}
	if req.WarehouseID != "" {
		db = db.Where("warehouse_id = ?", req.WarehouseID)
	}

	db.Count(&total)

	if pagination.Page > 0 && pagination.Size > 0 {
		offset := (pagination.Page - 1) * pagination.Size
		db = db.Limit(pagination.Size).Offset(offset)
	}

	err := db.Find(&transactions).Error
	return transactions, total, err
}
