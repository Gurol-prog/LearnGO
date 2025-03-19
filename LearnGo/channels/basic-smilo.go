package channels

func Cehf(Foods chan string) {

	food1 := "Yemek pişti"

	Foods <- food1

}

func Waiter(Foodwaiter chan string) {

	waiter1 := "Yemek dagıtıldı"

	Foodwaiter <- waiter1

}
