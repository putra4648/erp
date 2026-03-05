package dto

type WarehouseFindAllRequest struct {
	Name string `json:"name"`
	Page int    `json:"page"`
	Size int    `json:"size"`
}

type WarehouseDto struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}
