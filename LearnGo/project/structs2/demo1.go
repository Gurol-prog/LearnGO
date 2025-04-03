package structs2

import "fmt"

type producte struct {
	ID          int
	Name        string
	Description string
}

var U producte

func Deo1() {
	U.ID = 1
	U.Name = "Bilgisayar"
	U.Description = "Bişiler yapar"

	fmt.Printf("ID: %d\nName: %s\nAçıklama: %s\n",
		U.ID, U.Description, U.Name)
}
