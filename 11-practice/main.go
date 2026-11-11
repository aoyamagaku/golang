// func divide(a, b, c int) (int, int) {
// 	min := a
// 	max := a
// 	if b < min {
// 		min = b
// 	}
// 	if b > max {
// 		max = b
// 	}
// 	if c < min {
// 		min = c
// 	}
// 	if c > max {
// 		max = c
// 	}
// 	return min, max
// }
// return min(a, b, c), max(a, b, c) <-- решение в 1 строку

//	func fizzBuzz() {
//		for i := 1; i <= 30; i++ {
//			switch {
//			case i%3 == 0 && i%5 == 0:
//				fmt.Println("FizzBuzz")
//			case i%3 == 0:
//				fmt.Println("Fizz")
//			case i%5 == 0:
//				fmt.Println("Buzz")
//			default:
//				fmt.Println(i)
//			}
//		}
//	}
package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {
		if isEven(i) {
			fmt.Println(i)
		}
	}
}

func isEven(a int) bool {
	return a%2 == 0
}
