package main

import (
	"fmt"
	"strconv"
)

func main() {
	//TODO kullanıcı veri oluşturulacak
	fmt.Println("Kullanıcı oluşturma v1")
	user1 := &User{
		ID:        1,
		Firstname: "Oğğuz",
		Lastname:  "Cihan",
		Username:  "ocihan",
		Age:       24,
		Pay: &Payment{
			Salary: 3555,
			Bonus:  850.50,
		},
	}
	//fmt.Println(user1)
	fmt.Println(user1.GetFullName())
	fmt.Println(user1.GetPayment())
	fmt.Println("Maaş" + strconv.FormatFloat(user1.GetPayment(), 'g', -1, 64))

	fmt.Println("Kullanıcı oluşturma v2")
	user2 := NewUser()
	user2.Firstname = "Oğuz"
	user2.Lastname = "Cihan"
	user2.Pay = &Payment{Salary: 100, Bonus: 5}
	fmt.Println(user2.GetFullName())
	fmt.Println(user2.GetPayment())
	fmt.Println(user2.GetUserName())

}

//Kullanıcı yapısı
type User struct {
	ID        int
	Firstname string
	Lastname  string
	Username  string
	Age       int
	Pay       *Payment
}

//User Constructor
func NewUser() *User {
	u := new(User) //nesne üretildi
	u.Pay = NewPayment()
	return u
}

//Ödeme yapısı
type Payment struct {
	Salary float64
	Bonus  float64
}

//Payment Constructor
func NewPayment() *Payment {
	p := new(Payment)
	return p
}

//Methods
func (u User) GetFullName() string {
	//Ad soyad birleşitirip geriye döner
	return u.Firstname + " " + u.Lastname
}

func (u *User) GetUserName() string {
	return u.Username
}

func (u *User) GetPayment() float64 {
	pay := u.Pay.Salary + u.Pay.Bonus
	return pay

}
