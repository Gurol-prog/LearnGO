package maps

import "fmt"

func Demo1() {

	//key-value
	sozluk := make(map[string]string)
	sozluk["table"] = "Masa"
	sozluk["book"] = "kitap"
	sozluk["pencil"] = "kalem"
	sozluk["key"] = "Anahtar"

	fmt.Println(sozluk["book"])
	fmt.Println("Eleman Sayısı : ", len(sozluk))
	delete(sozluk, "book")
	fmt.Println("Eleman Sayısı : ", len(sozluk))

	deger, varMi := sozluk["lamb"]
	fmt.Println(deger)
	fmt.Println("Listede olma durumu : ", varMi)
}
