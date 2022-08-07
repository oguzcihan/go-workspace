package utils

type RestError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func BadRequest(_message string) *RestError {
	return &RestError{
		Message: _message,
		Status:  400,
		Error:   "Bad Request",
	}
}

func NotFound(_message string) *RestError {
	return &RestError{
		Message: _message,
		Status:  404,
		Error:   "Not found",
	}
}

func InternalError(_message string) *RestError {
	return &RestError{
		Message: _message,
		Status:  500,
		Error:   "Internal Server Error",
	}
}
