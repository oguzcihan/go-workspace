package Loops

import "fmt"

func Workshop() {
	aklimdakiSayi := 80
	tahminEdilen := 0
	tahminSayisi := 0
	fmt.Println("Bir sayı tahmin ediniz")
	fmt.Scanln(&tahminEdilen) //input girişi
	//fmt.Println(tahminEdilen)

	for aklimdakiSayi != tahminEdilen {
		if tahminEdilen < aklimdakiSayi {
			fmt.Println("Sayıyı arttır")
			fmt.Scanln(&tahminEdilen)
			tahminSayisi += 1

		} else if tahminEdilen > aklimdakiSayi {
			fmt.Println("Sayıyı azalt")
			fmt.Scanln(&tahminEdilen)
			tahminSayisi += 1

		}
	}
	baseDurum := "" //empty karakter anlamına gelir
	if tahminSayisi > 0 && tahminSayisi <= 3 {
		baseDurum = "Süper"
	} else {
		baseDurum = "Orta"
	}
	if aklimdakiSayi == tahminEdilen {
		fmt.Printf("Tahmin doğru. Tahmin Sayısı:%v :%v", tahminSayisi, baseDurum)
	}
}
