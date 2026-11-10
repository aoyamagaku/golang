package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

//"math/rand"
//"time"

// 🦆
// 🔋
// 🐍

func main() {
	// fmt.Println("Hello World!")
	// fmt.Println("Random number:", rand.Intn(5))
	// random := rand.Intn(21)
	// fmt.Println(random)
	// if random > 10 {
	// 	fmt.Println("Я сгенерировал число больше 10")
	// } else {
	// 	fmt.Println("Я балбес")
	// }

	// for i := 1; i <= 5; i++ {
	// 	fmt.Println("Hello World!", i)
	// 	time.Sleep(500 * time.Millisecond)
	// }

	score := 0

	fmt.Println("Get ready")
	fmt.Println("Score:", score)
	fmt.Println("")

	for {
		fmt.Println("You arrival at the pipe ")
		fmt.Println("🦆🔋🔋")
		fmt.Println("")

		fmt.Println("You fly through a pipe")
		fmt.Println("🔋🦆🔋")
		fmt.Println("")
		// 0 1 2 3 4
		if rand.IntN(5) == 1 {
			fmt.Println("You hit the pipe")
			fmt.Println("💀")
			fmt.Println("")
			fmt.Println("Game over")
			fmt.Println("Score:", score)
			fmt.Println("")
			break
		}

		fmt.Println("You flew through the pipe")
		fmt.Println("🔋🔋🦆")
		fmt.Println("")

		score++
		fmt.Println("Score:", score)
		fmt.Println("")

		time.Sleep(500 * time.Millisecond)
	}
}
