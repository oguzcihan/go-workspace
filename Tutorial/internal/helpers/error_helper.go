package helpers

import . "Tutorial/internal/models"

func SetError(err Error, message string) Error {
	err.IsError = true
	err.Message = message
	return err
}
