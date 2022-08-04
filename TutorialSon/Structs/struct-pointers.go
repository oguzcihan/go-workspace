package Structs

import "fmt"

type Vertex struct {
	X int
	Y int
}

func Struct() {
	v := Vertex{1, 3}
	p := &v
	p.X = 19 //(*p).X -> p.X aynı anlama gelmektedir
	fmt.Println(v)
}
