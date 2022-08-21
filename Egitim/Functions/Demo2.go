package Functions

import "fmt"

// variadic func :sınırsız değer atama
func ToplaVariadic(sayilar ...int) int {
	toplam := 0
	for i := 0; i < len(sayilar); i++ {
		toplam += sayilar[i]
	}
	return toplam
}

func Test() {
	sayilar := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(ToplaVariadic(sayilar...))
}
