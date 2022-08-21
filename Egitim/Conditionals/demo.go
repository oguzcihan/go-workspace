package Conditionals

import "fmt"

func Demo() {
	//üç adet int değişkeni
	//ekrana en büyük olanı yazdır
	var sayi1, sayi2, sayi3 int = 8, 20, 10
	//sayi1 = 5
	//sayi2 = 4
	//sayi3 = 10
	var max int = sayi1

	if sayi2 > max {
		max = sayi2
	}
	if sayi3 > max {
		max = sayi3
	}
	fmt.Println(max)

}
