package main

import (
	"fmt"
	"time"
)

func mine(transferPoint chan int, n int) {
	fmt.Println("Поход в шахту номер:", n, "начался")
	time.Sleep(1 * time.Second)
	fmt.Println("Поход в шахту номер:", n, "закончился")

	transferPoint <- 10
	fmt.Println("Поход номер", n, "уголь передан")
}

func main() {
	coal := 0

	transferPoint := make(chan int)

	initTime := time.Now()

	go mine(transferPoint, 1)
	go mine(transferPoint, 2)
	go mine(transferPoint, 3)

	coal += <-transferPoint
	coal += <-transferPoint
	coal += <-transferPoint

	fmt.Println("Добыли", coal)
	fmt.Println("Прошло время:", time.Since(initTime))

}
