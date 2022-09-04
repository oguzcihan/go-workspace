package utils

import (
	"github.com/go-playground/locales/tr"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	trTrans "github.com/go-playground/validator/v10/translations/tr"
	"log"
)

const (
	langTr = "tr"
	langEn = "en"
)

var (
	validate *validator.Validate
	//fieldError    string
	//ValidateError = NewError(fieldError, 404)
)

func LoadValidator() ut.Translator {
	translatorTr := tr.New()
	universal := ut.New(translatorTr, translatorTr)
	getTrans, found := universal.GetTranslator(langTr)
	if !found {
		log.Fatal("Translator bulunamadı")
	}
	validate = validator.New()
	err := trTrans.RegisterDefaultTranslations(validate, getTrans)
	if err != nil {
		return nil
	}
	return getTrans
}

func RequestValidate(model interface{}) error {

	getTrans := LoadValidator()
	var fieldError string

	err := validate.Struct(model)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, fe := range errs {
			fieldError = fe.Translate(getTrans)
		}
		return NewError(fieldError, 404)
	}
	return nil
}
