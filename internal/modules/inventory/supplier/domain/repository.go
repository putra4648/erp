package domain

import (
	"context"

	"github.com/google/uuid"
)

type SupplierRepository interface {
	Save(ctx context.Context, supplier *Supplier) error
	FindByID(ctx context.Context, id uuid.UUID) (*Supplier, error)
	FindAll(ctx context.Context, page, size int) ([]*Supplier, int64, error)
	Update(ctx context.Context, supplier *Supplier) error
	Delete(ctx context.Context, id uuid.UUID) error
}
