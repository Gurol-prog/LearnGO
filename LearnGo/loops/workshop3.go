package loops

import "fmt"

func WorkDemo3() {
	//220 ve 284 kardeş sayılardır
	sayi5 := 0
	sayi6 := 0
	sayi1k := 0
	sayi2k := 0

	fmt.Println("Kontrol edilecek sayılardan ilkini girin :")
	fmt.Scanln(&sayi5)
	fmt.Println("Kontrol edilecek sayılardan ikincisini girin :")
	fmt.Scanln(&sayi6)
	for i := 1; i < sayi5; i++ {
		if sayi5%i == 0 {
			sayi1k = sayi1k + i
		}

	}
	for i := 1; i < sayi6; i++ {
		if sayi6%i == 0 {
			sayi2k = sayi2k + i
		}

	}
	if (sayi1k == sayi6) && (sayi2k == sayi5) {
		fmt.Printf("%d sayisi ile %d sayısı kardeş sayılardır .", sayi5, sayi6)

	} else {
		fmt.Println("Bu sayılar kardeş değildir")
	}

}
