package myexmaple

import (
	"fmt"
	"time"
)

func SlowWorker(ch chan string) {
	time.Sleep(time.Second * 3)
	a := fmt.Sprintln("İş tamamlandı")
	ch <- a
}

func GoWorker() {

	ch := make(chan string)
	go SlowWorker(ch)

	select {
	case msg := <-ch:
		fmt.Println("Mesaj geldi : ", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("Çok bekledik vazgeç bu işten")

	}

}
