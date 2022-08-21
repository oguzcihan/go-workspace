package Channels

import (
	"fmt"
	"time"
)

func CiftSayilar(ciftSayiCn chan int) {
	toplam := 0
	for i := 0; i <= 10; i += 2 {
		//fmt.Println("Çift sayı", i)
		toplam += i
		fmt.Println("Çift sayı çalışıyor..")
		time.Sleep(1 * time.Second)
	}
	ciftSayiCn <- toplam //channelbu işlemin kendi yaşam döngüsünü bitirme anını yakalamak gerekiyor
}

func TekSayilar(tekSayiCn chan int) {
	toplam := 0
	for i := 1; i <= 10; i += 2 {
		//fmt.Println("Tek sayı", i)
		toplam += i
		fmt.Println("Tek sayı çalışıyor..")
		time.Sleep(1 * time.Second)
	}
	tekSayiCn <- toplam
}
