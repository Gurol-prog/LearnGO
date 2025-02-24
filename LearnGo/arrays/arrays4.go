package arrays

import "fmt"

func Demo4() {

	var sayilar [2][3]int

	// Kullanıcıdan veri al
	for i := 0; i < len(sayilar); i++ {
		for j := 0; j < len(sayilar[i]); j++ {
			fmt.Scanln(&sayilar[i][j])
		}
	}

	// Verileri ekrana yazdır
	for i := 0; i < len(sayilar); i++ {
		for j := 0; j < len(sayilar[i]); j++ {
			fmt.Print(sayilar[i][j], " ")
		}
		fmt.Println("") // Yeni satıra geç
	}
}
