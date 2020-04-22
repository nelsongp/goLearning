package main

import "fmt"

func main() {
	a := 10
	//en go las condicionales no se ponen entre parentesis
	if a > 5 {
		//se pueden crear variables dentro de los cuerpos de los ciclos qeu solo sirvan dentro del ciclo
		b := a / 2
		fmt.Println("a si bigger than 5 and b is ", b)
	} else {
		fmt.Println("a is less than or equal to 5")
		//b no es visible de este lado
	}
}
