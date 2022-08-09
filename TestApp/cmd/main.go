package main

import (
	"TestApp/internal/utils"
	"fmt"
	"github.com/go-playground/validator/v10"
)

func main() {

	//app.RunServer()

	//TODO:	//field eklenecek DONE
	//	//username tek
	//	//deleted sa tekrar oluşabilir username ile kontrol
	//	//lastname boş ise errror
	//	//email alanı gerçekten email mi?
	//	//tc kontrolü
	//	//zorunsuz alanlarda kullanılacak
	//	//password eklenecek validation kontrol
	//	//update delete olacak in progress

	p := &person{
		Name:  "sda",
		Email: "asdgf",
		Age:   0,
	}
	validate := validator.New()
	err := validate.Struct(p)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		fmt.Println("------ List of tag fields with error ---------")

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(utils.MessageForTag(err))
			//fmt.Println(err.StructField())
			//fmt.Println(err.ActualTag())
			//fmt.Println(err.Kind())
			//fmt.Println(err.Value())
			//fmt.Println(err.Param())
			fmt.Println("---------------")
		}
		return
	}
}

type person struct {
	Name  string `validate:"required,min=4,max=15" key:"name"`
	Email string `validate:"required,email"`
	Age   int    `validate:"required,numeric,min=18"`
}
