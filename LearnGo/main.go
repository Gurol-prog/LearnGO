package main

import (
	"fmt"
	"golesson/functions"
)

func main() {
	//variables.Demo1()
	//conditionals.Demo1()
	//conditionals.Demo2()
	//loops.Demo1()
	//loops.Demo2()
	//loops.WordkDemo()
	//loops.WorkDemo2()
	//loops.WorkDemo3()
	//arrays.Demo1()
	//arrays.Demo2()
	//arrays.Demo3()
	//arrays.Demo4()
	//slices.Demo2()
	//functions.SelamVer("Engin")
	//var sonuc = functions.Topla(5, 4)
	//fmt.Println(sonuc)
	// sonuc1, sonuc2, sonuc3, sonuc4 := functions.DortIslem(8, 9)
	// fmt.Println("Toplam : ", sonuc1)
	// fmt.Println("Çıkarım : ", sonuc2)
	// fmt.Println("Çarpım : ", sonuc3)
	// fmt.Println("Bölüm : ", sonuc4)
	var sonuc = functions.ToplaVariadic(1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 10, 11, 23, 15)
	fmt.Println(sonuc)

	sayilar := []int{4, 6, 7, 0, 11}
	fmt.Println(functions.ToplaVariadic(sayilar...))
}
