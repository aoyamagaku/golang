package main

import "fmt"

func main() {
	userArray := []User{
		User{
			Name:    "John",
			Rating:  4.5,
			Premium: true,
		},
		User{
			Name:    "Jane",
			Rating:  3.5,
			Premium: false,
		},
		User{
			Name:    "Mike",
			Rating:  1.2,
			Premium: true,
		},
		User{
			Name:    "Michael",
			Rating:  8.1,
			Premium: true,
		},
	}

	userArray = append(userArray, User{
		Name:    "Max",
		Rating:  7.9,
		Premium: false,
	})

	for i := range userArray {
		fmt.Println(userArray[i])
	}
	fmt.Println("------")
	//for i := range userArray { САМЫЙ СИНТАКСИЧЕСКИ ПРАВИЛЬНЫЙ ВАРИАНТ
	//	if userArray[i].Premium {
	//		userArray[i].Rating += 1
	//	}
	//}

	//for _, u := range userArray {  ТУТ ДОБАВЛЯЕТСЯ КОПИЯ, ОРИГИНАЛ НЕ МЕНЯЕТСЯ
	//	if u.Premium {
	//		u.Rating += 1
	//	}
	//}

	for i, _ := range userArray { // верный вариант, но синтаксис лучше как в самом первом
		if userArray[i].Premium {
			userArray[i].Rating += 1
		}
	}

	for i := range userArray {
		fmt.Println(userArray[i])
	}
}

type User struct {
	Name    string
	Rating  float64
	Premium bool
}
