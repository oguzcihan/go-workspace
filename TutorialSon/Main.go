package main

import (
	"GolangWorkspace/Structs"
	"fmt"
	"strconv"
)

/*
	Class yerine struct lar bulunur
*/
type Human struct {
	FirstName string
	LastName  string
	Age       int
}

//Default and blank constructor
//Method overload desteklemez
func NewHuman() *Human {
	//Constructor
	h := new(Human)
	return h
}

func NewHumanWithParams(firstname, lastname string, age int) *Human {
	h := new(Human)
	h.FirstName = firstname
	h.LastName = lastname
	h.Age = age
	return h
}

func main() {
	//TODO StructPointer
	Structs.Struct()

	//TODO farklı constuctor parametreli
	t := NewHumanWithParams("Oğuz", "Cihan", 24)
	var data = t.FirstName + " " + t.LastName + "/" + strconv.Itoa(t.Age)
	fmt.Println(data)

	//TODO constructor parametresiz
	a := NewHuman()
	a.Age = 5
	fmt.Println(a.Age)

	//TODO nesne örneği oluşturma
	x := Human{FirstName: "Cihan"}
	fmt.Println(x.FirstName)

	y := new(Human)
	y.FirstName = "Oğuz"
	fmt.Println(y.FirstName)

	//println("Hello World")

	//TODO package kullanımı
	//total := Utilities.Add(5, 6)
	//fmt.Println(total)

	//defer fmt.Println("world") //
	//fmt.Println("Hello")

	//TODO defer örneği
	//fmt.Printf("Connection open: %v\n", isConnected)
	//databaseProcessing()
	//fmt.Printf("Connection open: %v\n ", isConnected)
}

var isConnected bool = false

func databaseProcessing() {
	connect()
	fmt.Println("deferring disconnect")
	defer disconnect() //atlanacak
	fmt.Printf("connection open: %v\n", isConnected)
	fmt.Println("doing something")
}

func connect() {
	isConnected = true
	fmt.Println("Connected to database")
}
func disconnect() {
	isConnected = false
	fmt.Println("Disconnected")
}
