package DeferStatement

import "fmt"

type Product struct {
	productName string
	unitPrice   int
}

func (p Product) save() {
	fmt.Println("Ürün Kaydedildi", p.productName)
	defer Log() //loglar defer ile çağırılırsa rahat edilir çünkü log her zaman çalışması gerekir
}

func Log() {
	fmt.Println("Loglandı")
}

func Demo3() {
	p := Product{productName: "Laptop", unitPrice: 500}
	defer p.save()
	p = Product{productName: "Mouse", unitPrice: 50}
	fmt.Println("İşlem başarılı")
	fmt.Println(p.productName)
}
