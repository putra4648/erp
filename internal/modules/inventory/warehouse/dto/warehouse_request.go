package dto

type CreateWarehouseRequest struct {
	Name string `json:"name" validate:"required"`
	Code string `json:"code" validate:"required"`
}

type UpdateWarehouseRequest struct {
	Name string `json:"name" validate:"required"`
}
