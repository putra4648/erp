package service

type ProductError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *ProductError) Error() string {
	return e.Message
}

type UOMError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *UOMError) Error() string {
	return e.Message
}