package main

import "fmt"

func main() {
	foo()

	//funcion anonima
	func(x int) {
		fmt.Println("El significado de la vida es ", x)
	}(42)
}

func foo() {
	fmt.Println("Foo se ejecuto")
}
