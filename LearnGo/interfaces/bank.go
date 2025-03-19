package interfaces

import "fmt"

type CreditCalculator interface {
	Calculate() float64
}

type Mortgage struct {
	creditPaymentTotal float64
	address            string
	rate               float64
}

type Car struct {
	creditPaymentTotal float64
	carİnfo            string
	rate               float64
}

func (m Mortgage) Calculate() float64 {
	return m.rate * m.creditPaymentTotal / 12
}

func (c Car) Calculate() float64 {
	return c.rate * c.creditPaymentTotal / 12
}

func MonthlyPayment(credits []CreditCalculator) float64 {
	paymentTotal := 0.0

	for i := 0; i < len(credits); i++ {
		paymentTotal = paymentTotal + credits[i].Calculate()
	}
	return paymentTotal
}

func Demo3() {
	credit1 := Mortgage{rate: 10, creditPaymentTotal: 100000, address: "Ankara"}
	credit2 := Mortgage{rate: 12, creditPaymentTotal: 50000, address: "İstanbul"}
	credit3 := Car{rate: 15, creditPaymentTotal: 60000, carİnfo: "Honda"}

	credits := []CreditCalculator{credit1, credit2, credit3}

	total := MonthlyPayment(credits)

	fmt.Println("Toplam ödeme :", total)

}
