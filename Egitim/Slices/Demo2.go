package Slices

import "fmt"

func Demo2() {
	//slice kopyalama
	cities := []string{"Ankara", "İzmir", "İstanbul"} //make ile aynı işlemi yapar
	fmt.Println(cities)

	citiesCopy := make([]string, len(cities))
	fmt.Println(citiesCopy)
	copy(citiesCopy, cities)
	fmt.Println("Şehirler kopya", citiesCopy)

	cities = append(cities, "Adana")
	fmt.Println(cities)
	fmt.Println(citiesCopy)

	//filtre
	fmt.Println(cities[1:3]) //[x:y] y dahil değil
	fmt.Println(cities[1:])  //1 den başla sonuna kadar al
	fmt.Println(cities[:2])  //2 dahil olmadan alır
}
