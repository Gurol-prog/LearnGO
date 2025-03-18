package structs

import "fmt"

// ✅ Struct alanlarını büyük harfle başlattık (Dışarıdan erişim için)
type Car struct {
	Model    string
	Brand    string
	Year     int
	IsRented bool
}

// ✅ Fonksiyon adlarını büyük harfle yaptık (Başka paketlerden erişmek için)
func RentCar(c *Car) {
	if c.IsRented {
		fmt.Println(c.Model, "zaten kiralanmış!")
	} else {
		c.IsRented = true
		fmt.Println(c.Model, "başarıyla kiralandı.")
	}
}

func ReturnCar(c *Car) {
	if !c.IsRented {
		fmt.Println(c.Model, "zaten kiralanmamış.")
	} else {
		c.IsRented = false
		fmt.Println(c.Model, "başarıyla geri getirildi.")
	}
}
