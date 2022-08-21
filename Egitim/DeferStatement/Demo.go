package DeferStatement

import "fmt"

func A() {
	fmt.Println("A fonksiyonu çalıştı")
}

func B() {
	fmt.Println("B fonksiyonu çalıştı")
	defer A() //defer LIFO ile çalışır son giren ilk çıkar
	defer C()
}

func C() {
	fmt.Println("C fonksiyonu")
}
