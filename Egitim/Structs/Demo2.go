package Structs

import "fmt"

type Customer struct {
	Firstname string
	Lastname  string
	Age       int
}

// Struct metod ile ilişkilendirildiğinde buralar struct üzerinden çağırılır
func (c Customer) Save() {
	fmt.Println("Müşteri Eklendi:", c.Firstname)
}

func (c Customer) Update() {
	fmt.Println("Müşteri Eklendi:", c.Firstname)
}

func Demo2() {
	newCustomer := Customer{Firstname: "Oğuz", Lastname: "Cihan", Age: 24}
	newCustomer.Save()
	newCustomer.Update()
}
