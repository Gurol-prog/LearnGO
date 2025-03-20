package deferstatement

import "fmt"

func T(sayi int) string {
	defer DeferFunc()

	if sayi%2 == 0 {
		return "cift sayi"
	} else {
		return "tek sayi"
	}

}

func Test() {
	sonuc := T(9)

	fmt.Println(sonuc)

}

func DeferFunc() {

	fmt.Println("Bitti")

}
