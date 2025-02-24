package arrays

import "fmt"

func Demo3() {

	sayilar := [5]int{20, 30, 50, 10, 2}
	enbuyuk := sayilar[0]
	for i := 1; i < len(sayilar); i++ {
		if enbuyuk < sayilar[i] {
			enbuyuk = sayilar[i]
		}

	}
	fmt.Println(enbuyuk)
}
