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

var validate *validator.Validate

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

func RequestValidate(model interface{}) []string {
	getTrans := LoadValidator()
	var requestArray []string
	err := validate.Struct(model)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, fe := range errs {
			fieldError := fe.Translate(getTrans)
			requestArray = append(requestArray, fieldError)
		}
		return requestArray
	}
	return nil
}
