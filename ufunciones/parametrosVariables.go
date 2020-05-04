package main

import "fmt"

func main() {
	//funcion con argumentos variables
	xi := []int{2, 3, 4, 5, 6, 7, 8}

	//se le envia el slice de forma de array ... con los puntos asi se envia cada parametro
	//si se envia un slice hacia la funcion debe de ser el ultimo parametro, por regla del leguaje
	x := suma("hola", xi...)
	fmt.Println("La suma de los numeros enviados es de: ", x)
}

//funcion que recibe parametros variables
func suma(s string, x ...int) int {
	fmt.Println(s)
	fmt.Printf("%T\n", x)

	suma := 0
	for _, v := range x {
		suma += v
	}
	return suma
}
