package errors

type ErrorDto struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *ErrorDto) Error() string {
	return e.Message
}
