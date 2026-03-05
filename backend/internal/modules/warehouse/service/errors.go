package service

type WarehouseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *WarehouseError) Error() string {
	return e.Message
}
