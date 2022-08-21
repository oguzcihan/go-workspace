package Loops

import "fmt"

func Demo() {
	var text string = "Hello World"
	//for i := 0; i < ; i++ {
	//
	//}
	i := 1
	for i <= 5 {
		//infinite loop
		fmt.Println(text)
		i = i + 1
	}

	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}

}
