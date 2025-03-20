package deferstatement

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) save() {
	fmt.Println("Ürün kaydedildi: ", p.productName)
	defer Log()

}
func Demo3() {

	p := product{productName: "Laptop", unitPrice: 5000}

	defer p.save()

	fmt.Println("İslem basarılı")

}

func Log() {

	fmt.Println("Loglandı")

}
