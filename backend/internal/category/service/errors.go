package service

type CategoryError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *CategoryError) Error() string {
	return e.Message
}
