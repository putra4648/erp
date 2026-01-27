package services

import (
	"putra4648/erp/internal/models"
	"putra4648/erp/internal/repositories"

	"gorm.io/gorm"
)

type InventoryService struct {
	repository repositories.InventoryRepository
}

func NewInventoryService(db *gorm.DB) *InventoryService {
	return &InventoryService{
		repository: repositories.NewInventoryRepository(db),
	}
}

func (s *InventoryService) CreateStockAdjustmentResponse(stockAdjustment models.StockAdjustment) error {
	return s.repository.CreateStockAdjustment(stockAdjustment)
}
