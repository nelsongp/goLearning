package main

import "fmt"

func main() {
	defer foo()
	x, s := bar()
	fmt.Println(x, s)
	fmt.Println(sumaEntero([]int{1, 2, 3, 4, 5}...))
	fmt.Println(sumaSlice([]int{1, 2, 3, 4, 5}))
}

func bar() (int, string) {
	return 4, "Desde Bar"
}

func sumaEntero(x ...int) int {
	fmt.Println(x)
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func sumaSlice(x []int) int {
	fmt.Println(x)
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func foo() {
	fmt.Println("Retorno desde defer")
}
