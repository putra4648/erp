package dto

type CreateSupplierRequest struct {
	Name    string `json:"name" validate:"required"`
	Code    string `json:"code" validate:"required"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email" validate:"email"`
}

type UpdateSupplierRequest struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email" validate:"email"`
}

type SupplierFindAllRequest struct {
	Name string `json:"name"`
	Page int    `json:"page"`
	Size int    `json:"size"`
}