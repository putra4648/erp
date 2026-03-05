package service

type UOMError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *UOMError) Error() string {
	return e.Message
}
