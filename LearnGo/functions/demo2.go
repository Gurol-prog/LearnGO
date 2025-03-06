package functions

func DortIslem(sayi1 int, sayi2 int) (int, int, int, float32) {
	toplam := sayi1 + sayi2
	cikarim := sayi1 - sayi2
	carpim := sayi1 * sayi2
	bolum := sayi1 / sayi2
	return toplam, cikarim, carpim, float32(bolum)
}
