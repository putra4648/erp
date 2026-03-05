package dto

type SupplierFindAllRequest struct {
	Name string `json:"name"`
	Page int    `json:"page"`
	Size int    `json:"size"`
}

type SupplierDto struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Code    string `json:"code"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}
