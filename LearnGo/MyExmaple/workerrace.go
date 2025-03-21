package myexmaple

import (
	"fmt"
	"math/rand"
	"time"
)

func WorkerA(name string, ch chan string) {

	sayi := rand.Intn(5) + 1
	time.Sleep(time.Second * time.Duration(sayi))
	a := fmt.Sprintf("%s , İş bitti", name)
	ch <- a
}

func Race() {

	ch := make(chan string)

	go WorkerA("Worker 1", ch)
	go WorkerA("Worker 2", ch)
	go WorkerA("Worker 3", ch)

	msg := <-ch
	fmt.Println("🏁 Kazanan:", msg)

	for i := 1; i <= 3; i++ {
		select {
		case msg := <-ch:
			fmt.Println(msg, i)
		default:
			fmt.Println("Henüz veri yok, bekleniyor...")

		}

	}

}
