package structs

import (
	"fmt"
	"math/rand"
	"time"
)

type Character struct {
	Name        string
	Health      int
	AttackPower int
	IsAlive     bool
}

// 🗡 Saldırı Fonksiyonu
func Attack(attacker *Character, defender *Character) {
	if !attacker.IsAlive {
		fmt.Println("❌", attacker.Name, "öldüğü için saldırı yapamaz!")
		return
	}

	if !defender.IsAlive {
		fmt.Println("❌", defender.Name, "zaten ölü!")
		return
	}

	// Hasar ver
	damage := rand.Intn(attacker.AttackPower) + 1
	defender.Health -= damage

	fmt.Printf("⚔️ %s, %s'ye %d hasar verdi! (%d HP kaldı)\n",
		attacker.Name, defender.Name, damage, defender.Health)

	// Eğer canı 0 veya altına düşerse öldü
	if defender.Health <= 0 {
		defender.Health = 0
		defender.IsAlive = false
		fmt.Println("💀", defender.Name, "öldü!")
	}
}

// 📌 Karakter Durumunu Göster
func Status(c *Character) {
	if c.IsAlive {
		fmt.Printf("🏹 %s: %d HP (Yaşıyor!)\n", c.Name, c.Health)
	} else {
		fmt.Printf("💀 %s: %d HP (Ölü!)\n", c.Name, c.Health)
	}
}

func Battle() {

	rand.Seed(time.Now().UnixNano())

	// 🦸 Kahraman ve Canavar oluştur
	hero := Character{Name: "🌟 Kahraman", Health: 100, AttackPower: 20, IsAlive: true}
	monster := Character{Name: "👹 Canavar", Health: 80, AttackPower: 35, IsAlive: true}

	// 🛡 Karakterlerin başlangıç durumunu göster
	Status(&hero)
	Status(&monster)

	// ⚔️ Savaş başlasın!
	for hero.IsAlive && monster.IsAlive {
		Attack(&hero, &monster)
		if monster.IsAlive {
			Attack(&monster, &hero)
		}
	}

	// 🎯 Savaşın Sonucu
	fmt.Println("\n🎉 Savaş bitti!")
	Status(&hero)
	Status(&monster)

}
