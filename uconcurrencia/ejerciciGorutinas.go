package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//se debe de establecer el numeor de go rutinas que vamos a ejecutar
	wg.Add(2)

	go foo()
	go bar()

	fmt.Println("A punto de finalizar main.")
	wg.Wait()
}

func foo() {
	fmt.Println("HOla desde foo")
	wg.Done()
}

func bar() {
	fmt.Println("Hola desde bar")
	wg.Done()
}
