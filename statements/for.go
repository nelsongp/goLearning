package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		/* Podemos hacer un break en cualquier momento
		if i > 7 {
			break
		}*/
		fmt.Println(i)
	}

	b := 0
	for b < 10 {
		fmt.Println(b)
		b = b + 1
	}
	fmt.Println(b)

	//for range
	s := "Hello, world!"
	for k, v := range s {
		//se tiene que hacer un parse a string por que v si se imprime el solo es un rune, hay que pasarl o string
		fmt.Println(k, v, string(v))
	}
}
