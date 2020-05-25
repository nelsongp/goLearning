package main

import "fmt"

func main() {
	fmt.Println("La suma de es: ", miSuma(2, 4))
	fmt.Println("La suma de es: ", miSuma(1, 5))
	fmt.Println("La suma de es: ", miSuma(3, 9))

}

func miSuma(xi ...int) int {
	var sum int
	for _, v := range xi {
		sum += v
	}
	return sum
}
