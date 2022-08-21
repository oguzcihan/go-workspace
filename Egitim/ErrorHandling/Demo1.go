package ErrorHandling

import (
	"errors"
	"fmt"
)

func tahminEt(tahmin int) (string, error) {

	aklimdakiSayi := 50
	if tahmin < 1 || tahmin > 100 {
		return "", errors.New("1-100 arasında bir sayı giriniz")
	}

	if tahmin > aklimdakiSayi {
		return "Daha küçük sayı giriniz", nil
	}

	if tahmin < aklimdakiSayi {
		return "Daha küçük sayı giriniz", nil
	}
	return "Bildiniz", nil
}

func Demo1() {
	thm, _ := tahminEt(5)
	fmt.Println(thm)
	fmt.Println(tahminEt(101))
	fmt.Println(tahminEt(-10))

}
