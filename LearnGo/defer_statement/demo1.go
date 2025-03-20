package deferstatement

import "fmt"

func A() {
	fmt.Println("a foksiyonu çalıştı")

}
func C() {
	fmt.Println("c foksiyonu çalıştı")

}
func D() {
	fmt.Println("d foksiyonu çalıştı")

}

func B() {
	defer A()
	defer C()
	defer D()
	fmt.Println("b foksiyonu çalıştı")

}
