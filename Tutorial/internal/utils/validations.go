package utils

import (
	. "Tutorial/internal/dtos"
	"Tutorial/internal/langs"
	"github.com/go-playground/validator/v10"
)

func GenerateValidationResponse(err error) (response ValidationResponse) {
	response.Success = false

	var validations []Validation

	//get validation errors
	validationErrors := err.(validator.ValidationErrors)
	for _, value := range validationErrors {
		//get field&rule (tag)
		field, rule := value.Field, value.Tag

		//create validation object
		validation := Validation{Field: field(), Message: langs.GenerateValidationMessage(field(), rule())}

		//add validation object to validations
		validations = append(validations, validation)
	}

	//set validations response
	response.Validations = validations
	return response
}
