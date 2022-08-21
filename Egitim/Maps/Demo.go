package Maps

import "fmt"

func Demo() {
	//Dil desteği getirmek için kullanılabilir
	//key-value
	dict := make(map[string]string) //key ve value değerleri
	dict["table"] = "masa"
	dict["book"] = "kitap"
	dict["pencil"] = "kalem"

	fmt.Println(dict)
	fmt.Println(dict["book"])
	fmt.Println("eleman sayısı", len(dict))
	delete(dict, "book")
	fmt.Println("eleman sayısı", len(dict))
	fmt.Println(dict)

	deger, varMi := dict["table"]
	fmt.Println(deger)
	fmt.Println("Listede var mi?:", varMi)

	dict2 := map[string]string{"glass": "bardak", "camera": "kamera"}
	fmt.Println(dict2)
}
