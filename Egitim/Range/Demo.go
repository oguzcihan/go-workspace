package Range

import "fmt"

func Demo() {
	cities := []string{"Ankara", "Çorum"} //slice yapısı
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	//BUradaki i index anlamına gelir kullanmak istemezsek blank identifier kullanılır _
	//istenirse i değeri ve _ blank değeri kullanılmadan da yapılabilir
	for _, city := range cities {
		fmt.Println(city)
	}
}

func WorkShop() {
	//1-10 arasındaki sayıların tek olanları toplayan prog
	//for-range
	sayilar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tekToplam := 0
	for _, sayi := range sayilar {
		if sayi%2 != 0 {
			fmt.Println("Tek sayılar:", sayi)
			tekToplam += sayi
		}
	}
	fmt.Println("Tek sayı toplamı:", tekToplam)
}

func Demo2() {
	sozluk := map[string]string{"book": "kitap", "table": "masa"}
	for key, value := range sozluk {
		fmt.Println(key, value)
	}
}
