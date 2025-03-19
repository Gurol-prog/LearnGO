package channels

import (
	"fmt"
	"time"
)

// 📌 Şefin yemek hazırladığı fonksiyon
func Thef1(orderChannel chan string) {
	orders := []string{"Pizza", "Burger", "Pasta", "Salata", "Tatlı"}

	for _, order := range orders {
		fmt.Println("👨‍🍳 Şef hazırlıyor:", order)
		time.Sleep(time.Second) // Yemek hazırlamak zaman alıyor 🕒
		orderChannel <- order   // Yemek kanala gönderiliyor
	}

	close(orderChannel) // 📌 Tüm yemekler hazırlandı, kanal kapatılıyor
}

// 📌 Garsonun yemekleri aldığı fonksiyon
func Daiter(orderChannel chan string) {
	for order := range orderChannel {
		fmt.Println("🍽 Garson teslim etti:", order)
		time.Sleep(time.Second) // Teslim etmek zaman alıyor 🕒
	}
	fmt.Println("✅ Tüm siparişler teslim edildi!")
}

func Mulikas() {
	// 🔥 Channel oluştur
	orderChannel := make(chan string)

	// 🚀 Şefi ve garsonu çalıştır (goroutines ile)
	go Thef1(orderChannel)
	Daiter(orderChannel) // Garson çalışmaya başlıyor

	fmt.Println("🏁 Sipariş süreci tamamlandı!")
}
