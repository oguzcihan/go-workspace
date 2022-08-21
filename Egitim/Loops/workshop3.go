package Loops

import "fmt"

func Workshop3() {
	//220 ve 284 arkadaş sayı
	sayi1 := 220
	sayi2 := 284
	sayi1Toplam := 0
	sayi2Toplam := 0

	for i := 1; i <= sayi1; i++ {
		if sayi1%i == 0 {
			sayi1Toplam += i
		}
	}
	for i := 1; i <= sayi2; i++ {
		if sayi2%i == 0 {
			sayi2Toplam += i
		}
	}
	if sayi1Toplam == sayi2Toplam {
		fmt.Println("Arkadaş sayı")
	}
}
