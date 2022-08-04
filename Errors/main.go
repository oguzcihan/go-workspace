package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//file, err := os.Open("dosya.txt")
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//
	//fmt.Println(file.Name)

	//hata fırlatmak istendiğinde kullanılabilir
	//myError := errors.New("BU bir hata")
	//fmt.Println(myError)

	_, err := os.Open("abc.rar")
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err) //loglandı
		log.Fatal("Error", err)
		os.Exit(1)
	}
}
