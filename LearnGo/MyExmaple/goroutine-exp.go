package myexmaple

import (
	"fmt"
	"time"
)

func Worker(name string, ch chan string) {

	for i := 1; i <= 5; i++ {

		a := fmt.Sprintf("%s %d. mesaj", name, i)

		ch <- a

		time.Sleep(time.Second * 1)
	}
}

func ProjectWorker() {

	ch := make(chan string)

	go Worker("worker 1", ch)
	go Worker("worker 2", ch)
	go Worker("worker 3", ch)

	for i := 0; i < 15; i++ {

		select {
		case msg1 := <-ch:
			fmt.Println("ch1'den veri geldi:", msg1)
		default:
			fmt.Println("Henüz veri yok, bekleniyor...")

		}
	}

}
