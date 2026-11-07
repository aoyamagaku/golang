package main

import "fmt"

func main() {
	// subscribed := true

	// if subscribed {
	// 	fmt.Println("Thank you for subscribing on my channel")
	// } else {
	// 	fmt.Println("Please subscribe on my channel")
	// }

	score := 16
	fmt.Println(score)

	if score >= 8 && score <= 15 {
		fmt.Println("You HIT!")
	} else if score < 6 || score > 16 {
		fmt.Println("You are detected")
	} else {
		fmt.Println("You lose")
	}

	// if score > 15 {
	// 	fmt.Println("You are mega-cool")
	// } else if score > 10 {
	// 	fmt.Println("You are cool")
	// } else {
	// 	fmt.Println("You need to learn a lot")
	// }

	// switch {
	// case score > 15:
	// 	fmt.Println("You are mega-cool")
	// case score > 10:
	// 	fmt.Println("You are cool")
	// default:
	// 	fmt.Println("You need to learn a lot")
	// }

}
