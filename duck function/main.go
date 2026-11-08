package main

import "fmt"

// 🦆
// 🔋

func main() {
	score := 0

	fmt.Println("Get ready")

	for i := 1; i < 5; i++ {
		fmt.Println("You arrival at the pipe", i)
		fmt.Println("🦆🔋🔋")
		fmt.Println("")

		fmt.Println("You fly through a pipe", i)
		fmt.Println("🔋🦆🔋")
		fmt.Println("")

		fmt.Println("You flew through the pipe", i)
		fmt.Println("🔋🔋🦆")
		fmt.Println("")

		score++
		fmt.Println("Score:", score)
		fmt.Println("")
	}
}
