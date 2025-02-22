package loops

import "fmt"

// kullanıcıdan sayı girmesini iste
// kullanıcın giridği sayının asal olup olmadığın bulan örneği yazınız
func WorkDemo2() {

	sayi := 0

	fmt.Println("Lütfen bir sayı giriniz :")
	fmt.Scanln(&sayi)

	for i := 2; i < sayi; i++ {
		if sayi%i == 0 {
			fmt.Printf("%d sayisi asaldeğildir ", sayi)
			break
		} else {
			fmt.Printf("%d sayisi asaldır", sayi)
			break
		}
	}
}
