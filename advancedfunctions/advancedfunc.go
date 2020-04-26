package main

import "fmt"

func addOne(a int) int {
	return a + 2
}

//podemos tener funciones que reciban funciones
func printOperation(a int, f func(int) int) {
	fmt.Println(f(a))
}

func main() {
	// podemos asignarle una variable a la funcion sin mandarla a llamar no se le ponen() ya que estamos haciendo referencia a y no invocandola
	myAddOne := addOne
	fmt.Println(addOne(1))
	//acá llamamos a la funcoin y le pasamos un valor
	fmt.Println(myAddOne(1))

	//no se pueden declarar funciones dentro de otras funciones, se tienen que declarar anonimamente por medio de una variable
	myAddTwo := func(a int) int {
		return a + 1
	}
	fmt.Println(myAddTwo(1))

	printOperation(1, addOne)
}
