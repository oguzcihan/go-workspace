package Functions

func Topla(sayi1, sayi2 int) int {
	var toplam = sayi1 + sayi2
	return toplam
}

func DortIslem(sayi1, sayi2 int) (int, int, int, float32) {
	toplam := sayi1 + sayi2
	cikarim := sayi1 - sayi2
	carpim := sayi1 * sayi2
	bolum := float32(sayi1 / sayi2)

	return toplam, cikarim, carpim, bolum
}
