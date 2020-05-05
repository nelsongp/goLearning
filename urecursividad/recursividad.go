package main

import "fmt"

func main() {
	fmt.Println(4 * 3 * 2 * 1)
	n1 := factorial(8)
	fmt.Println(n1)
	n2 := cicloFact(8)
	fmt.Println(n2)
}

func factorial(n int) int {
	//caso base
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func cicloFact(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
