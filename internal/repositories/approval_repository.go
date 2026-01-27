package repositories

import (
	"gorm.io/gorm"
)

type ApprovalRepository struct {
	// Tambahkan field db di sini agar bisa digunakan
	db *gorm.DB
}

func NewApprovalRepository(db *gorm.DB) *ApprovalRepository {
	return &ApprovalRepository{
		db: db,
	}
}
