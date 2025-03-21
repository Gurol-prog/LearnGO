package errorhandling

import (
	"errors"
	"fmt"
)

func TahminEt(tahmin int) (string, error) {
	aklimdakiSayi := 61

	if tahmin > 100 || tahmin < 1 {
		return "", errors.New("1 dwen küçük veya 100den büyük sayı giriniz")
	} else if tahmin > aklimdakiSayi {

		return "Daha küçük bir sayi giriniz ", nil

	} else if tahmin < aklimdakiSayi {
		return "Daha büyük bir sayi giriniz ", nil
	} else {
		return "Tebrikler doğru sayıyı buldunuz", nil
	}

}

func Demo2(sayi int) {
	mesaj, hata := TahminEt(sayi)
	fmt.Println(mesaj, hata)
}
