package Interfaces

import "fmt"

func dogrula(i interface{}) {
	sayi, ok := i.(int)
	fmt.Println(sayi, ok)
}

func Demo3() {
	var sayi1 interface{} = "aefas" //dönüşemediği için ok false gelir
	dogrula(sayi1)

	var sayi2 interface{} = 45
	dogrula(sayi2)
}
