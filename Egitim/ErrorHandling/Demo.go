package ErrorHandling

import (
	"fmt"
	"os"
)

// Type assertion err.(os.PathError)
func Demo() {
	file, err := os.Open("demo1.txt")
	//dosya bulunursa err nil olur
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			//type assertion
			//pErr varsa ok true döner yoksa false döner
			fmt.Println("dosya bulunamadı: ", pErr.Path)
			return
		} else {
			fmt.Println("Dosya bulunamadı")
			fmt.Println(err.Error())
		}

	} else {
		fmt.Println("Dosya ismi", file.Name())
	}
}
