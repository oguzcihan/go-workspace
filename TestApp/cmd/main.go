package main

import "TestApp/internal/app"

func main() {

	app.RunServer()
	//TODO:
	//	//deleted sa tekrar oluşabilir username ile kontrol
	//	//update delete olacak

	//err := errors.New("use of err.String() detected!")
	//s := err.Error()
	//fmt.Printf(
	//	"version: %s\ntypes: %T / %T\nstring value via err.Error(): %q\n",
	//	runtime.Version(), err, s, s)
}
