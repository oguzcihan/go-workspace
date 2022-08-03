package main

import (
	"fmt"
)

/*
	Bir değişkeni isimlendirirken küçük harfle başlanırsa private olur büyük harf ile başlanırsa public olacaktır.
	Variables
	Go da her variable kullanılmak zorundadır. Kullanılmayan değişken tanımlanması istenmez
*/
func main() {
	fmt.Println("Hello Go")

	/*
		var i float32
		var i int
		i = 42
		var j int = 27
		k := 30
		fmt.Printf("%v,%T", i, i)*/
	/*
		Reflection bu değişkenin önce tipini algılar
		var message string
		message = "Hello Go "
		var messageTr = "Merhaba Go"
		fmt.Println(message + messageTr)
	*/

	// var a, b, c int = 3, 4, 5
	// fmt.Println(a + b + c)

	// var a int
	// a=8
	// fmt.Println(a)
	// var b string
	// var c float64
	// var d bool

	// var c int = 60
	// var k, o string = "sf", "sdf"
	// var s, b = "asf", true
	// fmt.Println(k, o, c, s, b)

	//şeklinde kullanılabilir
	// a ,b :="asfd",true
	// message := "merhaba"
	// fmt.Print(message)

	//sona nokta koyulduğudna ondalıklı anlamınna gelir (floating point)
	var myFloat32 float32 = 42.
	myComplex := complex(3, 4)
	println(myFloat32)
	println(myComplex)

	//Eğer bir function dışında isek := operatörü kullanılmaz
	//const değişkenler başlangıştça değer alırlar sonrasında değiştirilemezler
	// content = "test" //hata verir
	// fmt.Print(content)
	fmt.Print(k, l)
}

//Global variables
const content = "Go uygulamasi"
const pi = 3.14

var a, b, c = "test", true, 1
var d string
var (
	k = "test"
	l = 1
)
