package main

import (
	"fmt"
	"strconv"
)

func main() {
	//TODO
	str1 := "tuhe quick red fox"
	str2 := "jumped over"
	aNumber := 42
	isTrue := true
	stringLength, _ := fmt.Println(str1, str2) //println atama yaparsak bize stringlerin length ini verir
	fmt.Println(len(str1))                     // len() fonksiyonu string length leri verir
	fmt.Println("String length", stringLength)
	fmt.Printf("Value of anumber %v\n", aNumber) //% ile başlayanlara yer tutucu(place holder) denir. %v number değerin burada olacağı anlmaına gelir
	fmt.Printf("Value of isTrue %v\n", isTrue)
	fmt.Printf("Value of anumber as float: %.2f\n", float64(aNumber))
	fmt.Printf("Data types: %T\n", aNumber)

	//TODO konsoldan veri girişi
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter text:")
	// str, _ := reader.ReadString('\n')
	// fmt.Println(str)

	//TODO blank identifier
	// ikiyebol, _ := test(16)
	// fmt.Println(ikiyebol)

	// ikiyebol, ikiyebol2 := test(16)
	// fmt.Println(ikiyebol,ikiyebol2)

	//TODO Enum oluşturdu türü stringde olsa normalde gitse kabul ediyor
	printBrand(Lenovo)
	printBrand("Lenovo")

	//TODO Tür dönüşümü
	convert()
}

//TODO blank identifier
// func test(girdi int) (int, int) {
// 	islem1 := girdi / 2
// 	islem2 := girdi / 4

// 	return islem1, islem2
// }

//TODO Enum oluşturma
type Brand string

const (
	//Büyük küçük harfe duyarlı
	Lenovo  Brand = "Lenovo"
	Monster Brand = "Monster"
	Acer    Brand = "Acer"
)

func printBrand(brand Brand) {
	fmt.Println(brand)
}

//TODO Tür dönüşümleri
func convert() {
	var myString = "1"
	// var x = 10
	// var f float32 = 2.0

	// fmt.Println(myString, x, f) //yazdırıldı
	//string -> int dönüşümü
	//atoi fonksiyonu dışarıdan string alır geriye int döner eğer hata alırsa ikinci parametre olarak alır. Hatayı kullanmak istersek ikinci parametre olarak veririrz. Hata çıkmasın dersek blank identifier ile boş geçebiliriz
	number, _ := strconv.Atoi(myString)
	fmt.Println(number + 2)
	//Itoa int alır geriye string döner
	fmt.Println("Sonuc:" + strconv.Itoa(number))

	//Casting
	var i = 55
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Println(i, f, u)

	numTerms, sum := add(4, 5, 8, 7, 9, 7, 5, 1, 5)
	println("Added: ", numTerms, "terms for a total of", sum)

	fmt.Println(split(17))
}

func add(terms ...int) (int, int) {
	result := 0
	for _, term := range terms {
		result += term
	}
	return len(terms), result
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

//TODO
//defer
