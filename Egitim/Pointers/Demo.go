package Pointers

import "fmt"

func Demo(sayi *int) {
	*sayi = *sayi + 1
	fmt.Println("Demo daki sayı:", *sayi)
}
