package StringFunctions

import (
	"fmt"
	s "strings"
)

// case sensitive
func Demo() {
	name := "Oğuz"
	fmt.Println(s.Count(name, "O"))
	fmt.Println(s.Contains(name, "ğ")) //bir string içinde aranılan değer var mı yok mu?
	fmt.Println(s.Index(name, "o"))    //aranan ifadenin ilk kaçıncı index te olduğunu döndürür yoksa -1 döner
	fmt.Println(s.ToLower(name))
	fmt.Println(s.ToUpper(name))

}
