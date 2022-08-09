package utils

import "github.com/go-playground/validator/v10"

func MessageForTag(field validator.FieldError) string {
	switch field.Tag() {
	case "email":
		return "Geçersiz email"

	}
	return field.Error()
}
