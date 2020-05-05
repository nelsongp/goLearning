package main

import "fmt"

func main() {
	r := foo(bar, []int{1, 2, 3, 4, 9})
	fmt.Println(r)
}

func bar(xi []int) int {
	l := len(xi)
	if l == 0 {
		return 0
	} else if l == 1 {
		return xi[0]
	} else {
		return xi[0] + xi[l-1]
	}

}

func foo(f func(x []int) int, y []int) int {
	n := f(y)
	n++
	return n
}
