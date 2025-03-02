package slices

import "fmt"

func Demo1() {

	isimler := make([]string, 3)

	fmt.Println(isimler)
	isimler[0] = "Engin"
	isimler[1] = "Engin2"
	isimler[2] = "Engin3"
	isimler = append(isimler, "engin4")
	fmt.Println(isimler)
	fmt.Println(len(isimler))

}
