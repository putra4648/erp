package domain

import (
	"context"

	"github.com/google/uuid"
)

type WarehouseRepository interface {
	Save(ctx context.Context, warehouse *Warehouse) error
	FindByID(ctx context.Context, id uuid.UUID) (*Warehouse, error)
	FindAll(ctx context.Context, page, size int) ([]*Warehouse, int64, error)
	Update(ctx context.Context, warehouse *Warehouse) error
	Delete(ctx context.Context, id uuid.UUID) error
}
