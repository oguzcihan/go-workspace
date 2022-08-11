package utils

import (
	"fmt"
	"github.com/go-playground/locales/tr"
	. "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"log"
)

const (
	langTr = "tr"
	langEn = "en"
)

func RequestValidator(request interface{}) []string {

	getTrans := ToErrorResponse()
	validate := validator.New()

	_ = validate.RegisterTranslation("required", getTrans, func(ut Translator) error {
		return ut.Add("required", "{0} zorunlu alan", true) // see universal-translator for details
	}, func(ut Translator, fieldError validator.FieldError) string {
		retVal, _ := ut.T("required", fieldError.Field())
		return retVal
	})
	var err2 []string
	err := validate.Struct(request)
	for _, e := range err.(validator.ValidationErrors) {
		fieldError := e.Translate(getTrans)
		fmt.Println(e.Translate(getTrans))
		err2 = append(err2, fieldError)
		return err2
	}
	return nil
}

func ToErrorResponse() Translator {
	translator := tr.New()
	universal := New(translator, translator)
	getTrans, found := universal.GetTranslator(langTr)
	if !found {
		log.Fatal("Translator bulunamadı")
	}
	return getTrans
}
