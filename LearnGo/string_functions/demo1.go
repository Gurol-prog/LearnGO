package stringfunctions

import (
	"fmt"
	s "strings"
)

func Demo1() {

	isim := "engin"
	fmt.Println(s.Count(isim, "n"))
	fmt.Println(s.Contains(isim, "n"))
	fmt.Println(s.Index(isim, "n"))

	sonuc := s.Index(isim, "a")

	if sonuc != -1 {
		fmt.Println("A Harfi Var")
	} else {
		fmt.Println("A Harfi yok ")
	}
	fmt.Println(s.ToLower(isim))
	fmt.Println(s.ToUpper(isim))

}
