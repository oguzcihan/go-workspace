package utils

type CustomError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

//func NewCustomError(cError CustomError) *ErrorService {
//	return &ErrorService{err: cError}
//}
//
//type ErrorService struct {
//	err CustomError
//}

func BadRequest(_message string) *CustomError {
	return &CustomError{
		Message: _message,
		Status:  400,
	}
}

func NotFound(_message string) *CustomError {
	return &CustomError{
		Message: _message,
		Status:  404,
	}
}

func InternalError(_message string) *CustomError {
	return &CustomError{
		Message: _message,
		Status:  500,
	}
}
