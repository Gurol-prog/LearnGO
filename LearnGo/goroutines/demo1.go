package goroutines

import (
	"fmt"
	"time"
)

func CiftSayilar() {
	cifttSayilar := make([]int, 0)
	for i := 0; i <= 10; i += 2 {
		cifttSayilar = append(cifttSayilar, i)
		fmt.Println(i)
		time.Sleep(1 * time.Second)

	}
	fmt.Println("Çift sayılar : ", cifttSayilar)
}
func TekSayilar() {
	tektSayilar := make([]int, 0)
	for i := 1; i <= 10; i += 2 {
		tektSayilar = append(tektSayilar, i)
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Tek sayılar : ", tektSayilar)
}
