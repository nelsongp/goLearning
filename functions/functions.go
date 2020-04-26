package main

import "fmt"

//esta funcion retorna un parametro int, se declara el tipo pde parametro a retornar despues de los parentesis y antes de la lalve
//un funcion normal se declara func addNumbers(a int, b int){fmt.Println(a + b)}
func addNumbers(a int, b int) int {
	return a + b
}

//Go puede retornar mas de un valor por funciones declaradas
func divAndRemainder(a int, b int) (int, int) {
	return a / b, a % b
}

func main() {
	div, remainder := divAndRemainder(2, 3)
	fmt.Println(div, remainder)

	//no es necesario enviar los dos parametros para que la funcion retorne los valores
	div, _ = divAndRemainder(10, 4)
	fmt.Println(div)

	_, remainder = divAndRemainder(100, -100)
	fmt.Println(remainder)

	a := addNumbers(2, 3)
	fmt.Println(a)
}
