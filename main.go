package main

import "fmt"

func main() {
	score := 11

	// if score > 10 {
	// 	if score > 15 {
	// 		fmt.Println("Ты мега-крутой")
	// 	} else {
	// 		fmt.Println("Ты крутой")
	// 	}
	// } else {
	// 	fmt.Println("Тебе нужно многому научиться")
	// }

	switch {
	case score > 15:
		fmt.Println("You are mega-cool")
	case score > 10:
		fmt.Println("You are cool")
	default:
		fmt.Println("You need to learn a lot")
	}

}
