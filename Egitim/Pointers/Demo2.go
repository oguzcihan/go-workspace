package Pointers

import "fmt"

func Demo2(sayilar []int) {
	sayilar[0] = 100
	fmt.Println("Demodaki sayı", sayilar[0])
}
