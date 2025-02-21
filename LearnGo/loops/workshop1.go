package loops

import "fmt"

// 1'den 100'e kadar olan sayı tahmini oyunu
func WordkDemo1() {

	aklimdakiSayi := 80
	tahminEdilenSayi := 0
	hak := 5

	fmt.Println("Bir sayı tahmin ediniz (1-100 arasında):")
	fmt.Scanln(&tahminEdilenSayi)

	// **Geçersiz girişleri kontrol et**
	for tahminEdilenSayi < 1 || tahminEdilenSayi > 100 {
		fmt.Println("Geçersiz giriş! Lütfen 1 ile 100 arasında bir sayı giriniz:")
		fmt.Scanln(&tahminEdilenSayi)
	}

	for hak > 1 {
		if tahminEdilenSayi == aklimdakiSayi {
			hak--
			if hak == 4 {
				fmt.Println("Harika! İlk seferde bildiniz.")
			}
			fmt.Printf("Tahminiz doğru! Aklımdaki sayı: %d\n", aklimdakiSayi)
			break
		} else if tahminEdilenSayi < aklimdakiSayi {
			hak--
			fmt.Printf("Kalan hakkınız: %d\n", hak)
			fmt.Println("Daha yüksek bir sayı giriniz:")
			fmt.Scanln(&tahminEdilenSayi)

			// **Geçersiz giriş kontrolü tekrar**
			for tahminEdilenSayi < 1 || tahminEdilenSayi > 100 {
				fmt.Println("Geçersiz giriş! Lütfen 1 ile 100 arasında bir sayı giriniz:")
				fmt.Scanln(&tahminEdilenSayi)
			}

		} else if tahminEdilenSayi > aklimdakiSayi {
			hak--
			fmt.Printf("Kalan hakkınız: %d\n", hak)
			fmt.Println("Daha küçük bir sayı giriniz:")
			fmt.Scanln(&tahminEdilenSayi)

			// **Geçersiz giriş kontrolü tekrar**
			for tahminEdilenSayi < 1 || tahminEdilenSayi > 100 {
				fmt.Println("Geçersiz giriş! Lütfen 1 ile 100 arasında bir sayı giriniz:")
				fmt.Scanln(&tahminEdilenSayi)
			}
		}
	}

	if hak == 0 {
		fmt.Printf("Başarısız Tahmin edemediniz. Aklımdaki sayı: %d\n", aklimdakiSayi)
	}
}
