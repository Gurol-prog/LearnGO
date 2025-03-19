package interfaces

import "fmt"

// 📌 Ödeme Metodu Interface'i
type PaymentMethod interface {
	Pay(amount float64)
}

// 💳 Kredi Kartı Yapısı
type CreditCard struct {
	CardNumber string
}

func (cc CreditCard) Pay(amount float64) {
	fmt.Printf("💳 Kredi Kartı ile %.2f TL ödendi. Kart: **** **** **** %s\n", amount, cc.CardNumber[len(cc.CardNumber)-4:])
}

// 🏦 PayPal Yapısı
type PayPal struct {
	Email string
}

func (pp PayPal) Pay(amount float64) {
	fmt.Printf("📧 PayPal (%s) ile %.2f TL ödendi.\n", pp.Email, amount)
}

// ₿ Kripto Para Yapısı
type Crypto struct {
	WalletAddress string
}

func (c Crypto) Pay(amount float64) {
	fmt.Printf("₿ Kripto para cüzdanı (%s) ile %.2f TL ödendi.\n", c.WalletAddress[:6], amount)
}

func Mainse() {
	// 🔥 Ödeme yöntemlerini oluştur
	creditCard := CreditCard{CardNumber: "1234567812345678"}
	paypal := PayPal{Email: "user@example.com"}
	crypto := Crypto{WalletAddress: "3J98t1WpEZ73CNmQviecrnyiWrnqRhWNLy"}

	// 🛒 Kullanıcıdan ödeme yöntemini seçmesini iste
	var amount float64 = 250.75
	var method PaymentMethod

	fmt.Println("Ödeme yöntemi seçin: 1) Kredi Kartı 2) PayPal 3) Kripto")
	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		method = creditCard
	case 2:
		method = paypal
	case 3:
		method = crypto
	default:
		fmt.Println("❌ Geçersiz seçim!")
		return
	}

	// 💰 Ödemeyi gerçekleştir
	method.Pay(amount)
}
