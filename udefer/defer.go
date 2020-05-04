package main

import "fmt"

func main() {
	//Defer ejecucion es diferida para el momento en el que la funcion donde esta contenida retorna
	//pospone la ejecucion de la funcion hasta el final, primero hace bar y al terminar todo corre defer
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
