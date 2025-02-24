package arrays

import "fmt"

func Demo2() {

	var sehirler [5]string

	sehirler[0] = "ankara"
	sehirler[1] = "İstanbul"
	sehirler[2] = "İzmir"
	sehirler[3] = "Adana"
	sehirler[4] = "Diyarbakır"

	for i := 0; i < 5; i++ {
		fmt.Println(sehirler[i])
	}

}
