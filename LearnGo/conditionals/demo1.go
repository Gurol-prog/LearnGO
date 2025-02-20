package conditionals

import "fmt"

func Demo1() {

	var hesap float32 = 1000
	var cekilmekistenen float32 = 900

	if cekilmekistenen > hesap {

		fmt.Print("Hesaptaki Para Yetersiz")

	} else if cekilmekistenen < hesap {

		hesap = hesap - cekilmekistenen
		fmt.Printf("%.2f TL Para Çekildi , Kalan bakiyeniz : %.2f", cekilmekistenen, hesap)
	} else {
		fmt.Print("Tüm hesap bakiyesi çekildi")
	}

}
