package stringfunctions

import (
	"fmt"
	s "strings"
)

func Demo2() {

	isim := "Engin"
	fmt.Println(s.HasPrefix(isim, "Eng"))
	fmt.Println(s.HasSuffix(isim, "in"))
	fmt.Println(s.Index(isim, "gi"))

	harfler := []string{"e", "n", "g", "i", "n"}

	sonuc := s.Join(harfler, "*")
	fmt.Println(sonuc)
	//-1 hepsini değiştirir + lara geçersek şu kadar değiştir deriz
	fmt.Print(s.Replace(sonuc, "*", "+", -1))
	fmt.Println(s.Split(isim, "*"))
	fmt.Println(s.Repeat(sonuc, 5))

}
