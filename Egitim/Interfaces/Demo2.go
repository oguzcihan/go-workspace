package Interfaces

import "fmt"

// Polimorfizm kullanımı go dilinde bu şekilde yapılır örneği incele

// Mortgage:Ev kredisi
type Mortgage struct {
	creditPaymentTotal float64
	address            string
	rate               float64
}

type Car struct {
	creditPaymentTotal float64
	carInfo            string
	rate               float64
}

type CreditCalculater interface {
	Calculate() float64
}

func (m Mortgage) Calculate() float64 {
	return m.creditPaymentTotal * m.rate / 100 / 12
}

func (c Car) Calculate() float64 {
	return c.creditPaymentTotal * c.rate / 100 / 12
}

// Aylık ödeme hesaplaması
func CalculateMonthlyPayment(credits []CreditCalculater) float64 {
	paymentTotal := 0.0
	for i := 0; i < len(credits); i++ {
		paymentTotal += credits[i].Calculate()
	}
	return paymentTotal
}

func Demo2() {
	credit1 := Mortgage{rate: 10, creditPaymentTotal: 10000, address: "afsgs"}
	credit2 := Mortgage{rate: 15, creditPaymentTotal: 50000, address: "afsgs"}
	credit3 := Car{rate: 15, creditPaymentTotal: 60000, carInfo: "Passat"}

	credits := []CreditCalculater{credit1, credit2, credit3}
	total := CalculateMonthlyPayment(credits)
	fmt.Println("Toplam Ödeme:", total)
}
