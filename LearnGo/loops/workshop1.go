package loops

import "fmt"

//1 DEN YÜZE KADAR OLAN SAYI

func WordkDemo1() {

	aklimdakiSayi := 80
	tahminEdilenSayi := 0

	fmt.Println("Bir sayı tahmin ediniz")
	fmt.Scanln(&tahminEdilenSayi)

	hak := 5

	for hak > 1 {
		if tahminEdilenSayi == aklimdakiSayi {
			hak--
			if hak == 4 {

				fmt.Println("Harika ilk Seferde bilidiniz")
			}

			fmt.Printf("Tahminiz doğru aklımdaki sayi %d ", aklimdakiSayi)
			break
		} else if tahminEdilenSayi < aklimdakiSayi {
			hak--
			fmt.Printf("Kalan hakkınız : %d\n", hak)
			fmt.Println("Daha yüksek bir sayı tunuz")

			fmt.Scanln(&tahminEdilenSayi)

		} else if tahminEdilenSayi > aklimdakiSayi {
			hak--
			fmt.Printf("Kalan hakkınız : %d\n", hak)
			fmt.Println("Daha Küçük bir sayı tunuz")

			fmt.Scanln(&tahminEdilenSayi)

		}

	}
	if hak == 0 {
		fmt.Printf("Başarız , tahmin edemediniz . Aklımdaki sayi : %d", aklimdakiSayi)
	}
}
