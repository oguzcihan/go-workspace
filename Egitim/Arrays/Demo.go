package Arrays

import "fmt"

func Demo() {
	var sayilar [5]int //5 elemandan oluşan sayı dizisi
	sayilar[2] = 50    //2 indexe değer atandı
	fmt.Println(sayilar)

	for i := 0; i < len(sayilar); i++ {
		fmt.Print(sayilar[i], " ")
	}

}

func Demo2() {
	var cities [5]string
	cities[0] = "ankara"
	cities[1] = "Çorum"
	cities[2] = "İstanbul"
	cities[3] = "İzmir"
	cities[4] = "Bursa"
	fmt.Println(cities)
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}
}

func Demo3() {
	//En büyük sayıyı bulur
	sayilar := [5]int{20, 2, 3, 4, 5}
	fmt.Println(sayilar)

	enBuyuk := sayilar[1]

	for i := 0; i < len(sayilar); i++ {
		if sayilar[i] > enBuyuk {
			enBuyuk = sayilar[i]
		}
		fmt.Println(sayilar[i])

	}
	fmt.Println("En büyük:", enBuyuk)
}
