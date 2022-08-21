package utils

type CustomError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (c CustomError) Error() string {
	//TODO implement me
	panic("implement me")
}

func NewError(_message string, code int) error {
	return CustomError{
		Message: _message,
		Status:  code,
	}
}
