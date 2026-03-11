package dto

type SupplierDTO struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Code    string `json:"code"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}
