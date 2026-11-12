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

// func main() {
// 	for i := 1; i <= 50; i++ {
// 		if isPrime(i) {
// 			fmt.Println(i)
// 		}
// 	}
// }

// func isPrime(n int) bool {
// 	if n < 2 {
// 		return false
// 	}
// 	for d := 2; d < n; d++ {
// 		if n%d == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func main() {
// 	fmt.Println(sumDigits(345))
// 	fmt.Println(sumDigits(9))
// 	fmt.Println(sumDigits(1000))
// }

// func sumDigits(n int) int {
// 	total := 0
// 	for n > 0 {
// 		total += n % 10
// 		n /= 10
// 	}
// 	return total
// }

// func main() {
// 	x := 42 // переменная
// 	y := 52
// 	fmt.Println(x) // оригинал
// 	reset(&x)      // & дает адрес х, чтоб функция могла поменять значение оригинала
// 	fmt.Println(x) // оригинальный х, теперь равен 0, после функции
// 	reset2(y)
// 	fmt.Println(y)
// }

// func reset(p *int) { // функция обнуляет оригинал
// 	*p = 0
// }

// func reset2(n int) {
// 	n = 0
// }

package main

import "fmt"

func main() {
	x := 3
	y := 8
	fmt.Println("Before *:", x, y)
	swap(x, y)
	fmt.Println("After *:", x, y)
}

func swap(a, b int) {
	num := a
	a = b
	b = num
}
