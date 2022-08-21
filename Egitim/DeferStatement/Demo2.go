package DeferStatement

import "fmt"

func Demo2(sayi int32) string {
	defer DeferFunc()
	if sayi%2 == 0 {
		return "Çift sayı"
	}
	if sayi%2 != 0 {
		return "Tek sayı"
	}
	return ""
}

func Test() {
	sonuc := Demo2(9)
	fmt.Println("Sonuç:", sonuc)
}

func DeferFunc() {
	fmt.Println("Bitti")
}
