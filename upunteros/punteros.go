package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	var b *int = &a
	fmt.Println(b)
	//valor de referencia, extrae el valor de memoria de b
	fmt.Println(*b)

	//directamente en el valor de memoria modificacmos el valor
	*b = 43
	fmt.Println(a)
}
