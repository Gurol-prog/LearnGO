package main

import "golesson/project/structs2"

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
	// var sonuc = functions.ToplaVariadic(1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 10, 11, 23, 15)
	// fmt.Println(sonuc)

	// sayilar := []int{4, 6, 7, 0, 11}
	// fmt.Println(functions.ToplaVariadic(sayilar...))
	//maps.Demo1()
	//rangex.Demo1()
	//rangex.Demo2()
	//rangex.Demo3()

	// sayi := 20
	// pointers.Demo1(&sayi)
	// fmt.Println("Maindeki sayi ", sayi)

	//sayilar := []int{1, 2, 3}
	//pointers.Demo2(sayilar)
	//fmt.Println("Maindeki sayilar : ", sayilar[0])
	//structs.Demo2()
	// Bir araç nesnesi oluştur
	// 	car1 := structs.Car{Model: "Sedan", Brand: "Honda", Year: 2025, IsRented: false}
	// 	structs.RentCar(&car1)
	// 	structs.RentCar(&car1)
	// book1 := structs.Book{Title: "Ateşten Gömlek", Author: "Halide edip adıvar", Year: 1923, IsBorrowed: false}
	// structs.BorrowBook(&book1)
	// structs.ReturnBook(&book1)
	// structs.BorrowBook(&book1)
	// structs.BorrowBook(&book1)
	// structs.ReturnBook(&book1)
	// structs.ReturnBook(&book1)

	// Rastgele sayı üretmek için

	// structs.Battle()

	// ciftSayiToplamCn := make(chan int)
	// tekSayiToplamCn := make(chan int)
	// go channels.CiftSayilar(ciftSayiToplamCn)
	// go channels.TekSayilar(tekSayiToplamCn)

	// ciftSayiToplam, tekSAyitoplam := <-ciftSayiToplamCn, <-tekSayiToplamCn

	// carpim := ciftSayiToplam * tekSAyitoplam

	// fmt.Println("Çarpım : ", carpim)

	// foods := make(chan string)
	// waiter := make(chan string)
	// go channels.Cehf(foods)
	// go channels.Waiter(waiter)

	// foods2, waiterss := <-foods, <-waiter

	// fmt.Println(foods2)
	// fmt.Println(waiterss)

	// channels.Mulikas()

	// interfaces.Demo1()

	// interfaces.Demo3()

	// interfaces.Demo4()

	// deferstatement.Demo3()

	// interfaces.Demo5()

	// for i := 0; i < 5; i++ {
	// 	tutulanSayi := 0
	// 	fmt.Scanln(&tutulanSayi)
	// 	errorhandling.Demo2(tutulanSayi)
	// }

	// fmt.Println(errorhandling.TahminEt2(102))
	// myexmaple.ProjectWorker()

	// myexmaple.RaceB()
	// case sensitive
	// stringfunctions.Demo2()
	// restfull.Demo2()

	// project.AddProduct()
	// project.GetAllProduct()

	// products, _ := project.GetAllProductRefactor()

	// for i := 0; i < len(products); i++ {
	// 	fmt.Println(products[i].ProductName)
	// }

	// var tasks []taskmanager.Task

	// taskmanager.AddTask(&tasks, "Go öğren", "Goroutines konusunu bitir")
	// taskmanager.AddTask(&tasks, "Proje geliştir", "Mini bir CLI uygulama yap")

	// fmt.Println("Tüm görevler:")
	// taskmanager.ListTasks(tasks)

	// taskmanager.MarkAsDone(&tasks, 1)

	// taskmanager.ListTasks(tasks)

	structs2.Deo1()
}
