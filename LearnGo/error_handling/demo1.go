package errorhandling

import (
	"fmt"
	"os"
)

// type assertion
func Demo1() {
	f, err := os.Open("demo2.txt")
	// olurda doysa bulunur ise erroun değeri nil dediğimzi bir değer olur
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Dosyabulunamadı", pErr.Path)
			return
		} else {
			fmt.Println("dosya bulunamadı,")
			return
		}

	} else {
		fmt.Println(f.Name())
	}
}
