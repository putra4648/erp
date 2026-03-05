package dto

type AdjustmentReasonRequest struct {
	Name        string `json:"name" validate:"required"`
	AccountCode string `json:"account_code" validate:"required"`
}

type AdjustmentReasonDto struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}
