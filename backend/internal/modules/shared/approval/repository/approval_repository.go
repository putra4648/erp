package repository

import (
	"context"
	"errors"
	"putra4648/erp/internal/modules/shared/approval/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ApprovalRepository interface {
	Create(ctx context.Context, docCode string, referenceID uuid.UUID) (domain.ApprovalTransaction, error)
}

type approvalRepository struct {
	db *gorm.DB
}

func NewApprovalRepository(db *gorm.DB) *approvalRepository {
	return &approvalRepository{db}
}

func (r *approvalRepository) Create(ctx context.Context, docCode string, referenceID uuid.UUID) (domain.ApprovalTransaction, error) {
	var res domain.ApprovalTransaction
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var workflow domain.ApprovalWorkflow
		if err := tx.Where("doc_code = ?", docCode).First(&workflow).Error; err != nil {
			return errors.New("workflow tidak ditemukan")
		}

		newApproval := domain.ApprovalTransaction{
			ID:          uuid.New(),
			WorkflowID:  workflow.ID,
			ReferenceID: referenceID,
			CurrentStep: 1,
			Status:      "PENDING",
		}

		if err := tx.Create(&newApproval).Error; err != nil {
			return err
		}

		res = newApproval
		return nil
	})

	return res, err
}
