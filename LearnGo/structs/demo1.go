package structs

import "fmt"

func Demo1() {
	fmt.Println(product{"Bilisauar", 5000, "xyz", 250})
	fmt.Println(product{name: "Bilisauar", discountRate: 5000, brandis: "xyz"})
}

type product struct {
	name         string
	unitPrice    float32
	brandis      string
	discountRate int
}
