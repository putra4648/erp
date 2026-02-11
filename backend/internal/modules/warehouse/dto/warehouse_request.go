package dto

type CreateWarehouseRequest struct {
	Name string `json:"name" validate:"required"`
	Code string `json:"code" validate:"required"`
}

type UpdateWarehouseRequest struct {
	Name string `json:"name" validate:"required"`
}

type WarehouseFindAllRequest struct {
	Name string `json:"name"`
	Page int    `json:"page"`
	Size int    `json:"size"`
}
