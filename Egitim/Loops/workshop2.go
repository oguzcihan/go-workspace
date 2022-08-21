package Loops

import "fmt"

func Workshop2() {
	//1.kullanıcıdan sayı al
	//2. 23 asaldır sayısı döner

	sayi := 0
	fmt.Println("Bir sayı giriniz")
	fmt.Scanln(&sayi)
	asalMi := true
	for i := 2; i < sayi; {
		if sayi%i == 0 {
			asalMi = false
		}
		if asalMi {
			fmt.Println("Asaldır")
			fmt.Scanln(&sayi)
		} else {
			fmt.Println("Asal değiildir")
			break
		}
	}

	//if asalMi {
	//	fmt.Println("Asaldır")
	//} else {
	//	fmt.Println("Asal değildir")
	//}
}
