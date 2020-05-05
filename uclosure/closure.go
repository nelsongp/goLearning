package main

import "fmt"

func main() {
	a := incrementor()
	b := incrementor()

	fmt.Println(a())
	//en cada corrida como se tiene una copia en memoria el valor de x ira incrementando
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func incrementor() func() int {
	var x int
	//closure se lleva una copia del entorno donde esta definido
	return func() int {
		x++
		return x
	}
}
