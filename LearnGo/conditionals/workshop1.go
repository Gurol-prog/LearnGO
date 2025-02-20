package conditionals

import "fmt"

func Demo2() {
	//üç adet int değişken tanımlayalınız.
	// Ekrana en büyük olanı yazdırınız.
	var sayi1 int = 10
	var sayi2 int = 20
	var sayi3 int = 30

	// sayi1 := 10
	// sayi2 := 20
	// sayi3 := 30
	// bu şekildede tanımlaya bilirdik

	if sayi1 > sayi2 {
		if sayi1 > sayi3 {
			fmt.Println("Sayi1 En büyük sayi :" + fmt.Sprintf("%d", sayi1))

		}
	} else if sayi2 > sayi3 {

		fmt.Println("Sayi2 En büyük sayi :" + fmt.Sprintf("%d", sayi2))
	} else {

		fmt.Println("Sayi3 En büyük sayi :" + fmt.Sprintf("%d", sayi3))
	}

}
