package structs

import "fmt"

type customeris struct {
	name  string
	lname string
	age   int
}

func (c customeris) addCustomer() {
	fmt.Println("eklendi", c.name)

}

func Demo2() {
	c := customeris{name: "Engin", lname: "Demiol", age: 35}

	c.addCustomer()

}
