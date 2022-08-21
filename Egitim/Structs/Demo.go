package Structs

import "fmt"

type Product struct {
	Name         string
	UnitPrice    float64
	Brand        string
	DiscountRate int
}

func Demo() {
	fmt.Println(Product{Name: "Bilgisayar", UnitPrice: 500, Brand: "xyz"})
	fmt.Println(Product{Name: "Bilgisayar", UnitPrice: 500})
}
