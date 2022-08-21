package ErrorHandling

import (
	"fmt"
)

type borderException struct {
	parameter int
	message   string
}

func (b *borderException) Error() string {
	//error bilgisi döndürüldü
	return fmt.Sprintf("%d--%s", b.parameter, b.message)
}

func Demo3(tahmin int) (string, error) {
	if tahmin < 1 || tahmin > 100 {
		return "", &borderException{parameter: 1, message: "sınırların dışında kaldı"}
	}
	return "bildiniz", nil
}
