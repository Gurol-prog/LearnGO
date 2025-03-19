package interfaces

import "fmt"

type Zoo interface {
	makeSound() string
}

type Dog struct {
	name string
}
type Cat struct {
	name string
}
type Bird struct {
	name string
}

func (d Dog) makeSound() string {

	return "Hav Hav"

}
func (c Cat) makeSound() string {

	return "Miyav"

}
func (b Bird) makeSound() string {

	return "Cik Cik"

}

func Animals(z Zoo) string {

	a := z.makeSound()
	return a

}
func Demo4() {
	animals1 := Dog{name: "Skuby"}
	animals2 := Cat{name: "Flapy"}
	animals3 := Bird{name: "Garfield"}

	a := Animals(animals1)
	fmt.Println(animals1.name, "sesi : ", a)
	Animals(animals2)
	Animals(animals3)

}
