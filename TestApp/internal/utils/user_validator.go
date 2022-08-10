package utils

import (
	"TestApp/internal/model"
	"fmt"
	"github.com/go-playground/locales/tr"
	uTranslator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"log"
)

func RegisterUserValidator(user *model.User) {
	//dto ile çekilmeli user
	translator := tr.New()
	universal := uTranslator.New(translator, translator)
	getTrans, found := universal.GetTranslator("tr")
	if !found {
		log.Fatal("Translator bulunamadı")
	}
	validate := validator.New()

	_ = validate.RegisterTranslation("required", getTrans, func(ut uTranslator.Translator) error {
		return ut.Add("required", "{0} zorunlu alan", true) // see universal-translator for details

	}, func(ut uTranslator.Translator, fieldError validator.FieldError) string {
		retVal, _ := ut.T("required", fieldError.Field())
		return retVal
	})
	err := validate.Struct(&user)
	for _, e := range err.(validator.ValidationErrors) {
		fmt.Println(e.Translate(getTrans))

	}
}
