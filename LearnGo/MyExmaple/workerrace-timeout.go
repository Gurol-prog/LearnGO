package myexmaple

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func workerB(ctx context.Context, name string, ch chan string) {
	sure := rand.Intn(6) + 1 // 1 ile 6 saniye arasında rastgele süre
	select {
	case <-time.After(time.Duration(sure) * time.Second): // İş süresi kadar bekle
		ch <- fmt.Sprintf("🏁 %s: İş bitti (%d saniye)", name, sure)
	case <-ctx.Done(): // Timeout süresi dolduysa
		ch <- fmt.Sprintf("⏳ DISKALIFIYE: %s süresini aştı!", name)
	}
}

func RaceB() {
	ch := make(chan string)
	// Timeout için context oluştur
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel() // İş bitince temizlemeyi unutma!

	go workerB(ctx, "Worker1", ch)
	go workerB(ctx, "Worker2", ch)
	go workerB(ctx, "Worker3", ch)
	go workerB(ctx, "Worker4", ch)
	go workerB(ctx, "Worker5", ch)

	for i := 1; i <= 5; i++ {

		msg := <-ch
		fmt.Println(msg)

	}
}
