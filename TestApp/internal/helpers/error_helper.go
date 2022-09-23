package helpers

import . "TestApp/internal/model"

func SetError(err Error, message string) Error {
	err.IsError = true
	err.Message = message
	return err
}
