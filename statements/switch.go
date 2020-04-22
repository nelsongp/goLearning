package main

import (
	"fmt"
	"os"
)

func main() {
	word := os.Args[1]
	//podemos declarar mas de una variable en el switch y usarla
	switch l := len(word); {
	case word == "hello":
		fmt.Println("hi yourself")
		//se se cuimple este esscenario hace este y el siguiente
		fallthrough
	//se pueden tener mas de un valor en el case
	case word == "bye":
		fmt.Println("So long!")
	case word == "farewell":
	case l == 1:
		fmt.Println("one letter")
	case word == "greetings":
		fmt.Println("Salutations!")
	default:
		fmt.Println("i don´t know, but length was ", l)
	}
}
