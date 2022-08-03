package main

func main() {

	// var b int = 15
	// var a int
	// numbers := [6]int{1, 2, 3, 4, 5}

	// for a := 0; a < 10; a++ {
	// 	fmt.Printf("value of a: %d\n", a)
	// }

	// for a < b {
	// 	a++
	// 	fmt.Printf("value of a: %d\n", a)
	// }

	// for i, x := range numbers {
	// 	fmt.Printf("value of x: %d at %d\n", x, i)
	// }

	//TODO Asal Sayı
	// var i, j int
	// for i = 2; i < 20; i++ {
	// 	for j = 2; j <= (i / j); j++ {
	// 		if i%j == 0 {
	// 			fmt.Printf("%d Asal değil\n", i)
	// 			break
	// 		}
	// 	}
	// 	if j > (i / j) {
	// 		fmt.Printf("%d Asal\n", i)
	// 	}
	// }

	//TODO LOOP ifadesi
	// 	var a int = 10
	// LOOP:
	// 	for a < 20 {
	// 		if a == 15 {
	//eğer a==15 ise loop a geri dön yapar ve 15 sayısını atlar
	// 			a = a + 1
	// 			goto LOOP
	// 		}
	// 		fmt.Printf("%d\n", a)
	// 		a++
	// 	}

	// var int write=max(5, 6)
}

//TODO functions
//Fonksiyon dönüş türü int
func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
