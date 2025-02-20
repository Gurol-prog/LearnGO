package variables

import "fmt"

func Demo1() {

	var metin string = "Merhaba Türkiye!"
	fmt.Print(metin)
	fmt.Print(metin)
	fmt.Print(metin)
	fmt.Print(metin)
	fmt.Println(metin)

	var kdv int = 10
	fmt.Println(kdv)
	fmt.Println(100 + (100 * kdv / 100))

	var kdv2 float32 = 10.5

	fmt.Println(kdv2)

	kdv3 := "deneme"
	fmt.Println(kdv3)

	kdv4 := 10

	fmt.Println(kdv4)

	// := kullanarak da değişken tanımlayabiliriz belirtmemize gerek yoktur

	var durum bool

	var metin1 string = "engin"
	var metin2 string = "Ahmet"

	durum = metin1 == metin2

	fmt.Println(durum)

}
