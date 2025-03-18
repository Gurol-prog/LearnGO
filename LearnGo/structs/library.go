package structs

import "fmt"

// Title (string) → Kitap adı
// Author (string) → Yazar
// Year (int) → Basım yılı
// IsBorrowed (bool) → Ödünç alındı mı?
// BorrowBook() Fonksiyonu

// Parametre olarak pointer *Book alacak.
// Kitap zaten ödünç alınmışsa mesaj versin, değilse ödünç alsın.
// ReturnBook() Fonksiyonu

// Parametre olarak pointer *Book alacak.
// Kitap zaten kütüphanede ise mesaj versin, değilse geri getirilsin.
// main() İçinde Kullanım

// Birkaç kitap oluştur.
// Ödünç alma ve geri getirme işlemlerini yap.

type Book struct {
	Title      string
	Author     string
	Year       int
	IsBorrowed bool
}

func BorrowBook(b *Book) {
	if b.IsBorrowed {
		fmt.Println(b.Title, " Adlı kitap zaten kiralanmış")
	} else {
		b.IsBorrowed = true
		fmt.Println(b.Title, " Adlı kitap kiralandı ")
	}

}

func ReturnBook(b *Book) {
	if !b.IsBorrowed {
		fmt.Println(b.Title, " Adlı kitap zaten kütüphanede")
	} else {
		b.IsBorrowed = false
		fmt.Println(b.Title, " Adlı kitap geri getirildi ")
	}

}
