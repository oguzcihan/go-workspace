package Interfaces

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	//p.r kare
	return math.Pi * c.Radius * c.Radius
}

func School(s Shape) {
	fmt.Println("Şeklin ALANI:", s.Area())

}

func Demo() {
	r := Rectangle{Width: 20, Height: 12}
	School(r)

	c := Circle{Radius: 20}
	School(c)
}
