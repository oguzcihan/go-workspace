package helpers

type CustomError struct {
	Message string `json:"message"`
	Status  int    `json:"code"`
	Err     error  `json:"-"`
}

func (c CustomError) Error() string {
	return c.Err.Error()
}

func NewError(_message string, code int) error {
	return CustomError{
		Message: _message,
		Status:  code,
	}
}
