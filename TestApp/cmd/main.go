package main

import "TestApp/internal/app"

func main() {

	app.RunServer()

	//TODO:	//field eklenecek DONE
	//	//username tek
	//	//deleted sa tekrar oluşabilir username ile kontrol
	//	//lastname boş ise errror
	//	//email alanı gerçekten email mi?
	//	//tc kontrolü
	//	//zorunsuz alanlarda kullanılacak
	//	//password eklenecek validation kontrol
	//	//update delete olacak in progress

	//TestUserValidator()
	//fmt.Println("Test:" + os.Getenv("PgConnection"))
}

//type User struct {
//	Email    string `json:"email" validate:"required,email"`
//	Name     string `json:"name" validate:"required"`
//	Password string `json:"password" validate:"passwd"`
//}
//
//func TestUserValidator() {
//	translator := en.New()
//	uni := ut.New(translator, translator)
//	// this is usually known or extracted from http 'Accept-Language' header
//	// also see uni.FindTranslator(...)
//	trans, found := uni.GetTranslator("en")
//	if !found {
//		log.Fatal("translator not found")
//	}
//
//	validate := validator.New()
//
//	//if err := en_translations.RegisterDefaultTranslations(v, trans); err != nil {
//	//	log.Fatal(err)
//	//}
//
//	_ = validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
//		return ut.Add("required", "{0} zorunlu alan", true) // see universal-translator for details
//	}, func(ut ut.Translator, fieldError validator.FieldError) string {
//		retVal, _ := ut.T("required", fieldError.Field())
//		return retVal
//	})
//
//	_ = validate.RegisterTranslation("email", trans, func(ut ut.Translator) error {
//		return ut.Add("email", "{0} geçersiz", true) // see universal-translator for details
//	}, func(ut ut.Translator, fe validator.FieldError) string {
//		t, _ := ut.T("email", fe.Field())
//		return t
//	})
//
//	_ = validate.RegisterTranslation("passwd", trans, func(ut ut.Translator) error {
//		return ut.Add("passwd", "{0} geçersiz", true) // see universal-translator for details
//	}, func(ut ut.Translator, fe validator.FieldError) string {
//		t, _ := ut.T("passwd", fe.Field())
//		return t
//	})
//
//	_ = validate.RegisterValidation("passwd", func(fl validator.FieldLevel) bool {
//		return len(fl.Field().String()) > 6
//	})
//
//	a := User{
//		Email:    "a",
//		Password: "1234",
//	}
//	err := validate.Struct(a)
//	//test := err.(validator.ValidationErrors)
//	//test.Translate(trans)
//	for _, e := range err.(validator.ValidationErrors) {
//		fmt.Println(e.Translate(trans))
//	}
//}
