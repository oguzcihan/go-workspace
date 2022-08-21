package Slices

import "fmt"

func Demo() {
	names := make([]string, 3)
	names[0] = "oğuz0"
	names[1] = "oğuz1"
	names[2] = "oğuz2"
	//names[3] = "oğuz2" bu şekilde hata verir ama append ile yaparsak bu slice genişler
	names = append(names, "oğuz3append")
	fmt.Println(names)
	fmt.Println(len(names))
}
