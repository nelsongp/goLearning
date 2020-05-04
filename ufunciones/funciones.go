package main

import "fmt"

func main() {
	foo()
	bar("James.")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	//funcion que retorna dos parametros
	x, y := saludar("Nelson", "guardado")
	fmt.Println(x, " ", y)

	//funcion con argumentos variables
	sum1 := sumaNum(2, 3, 4, 5, 6, 7, 8, 9)
	sum2 := sumaNum(2, 3, 4, 5)

	fmt.Println("La suma de los numeros enviados es de: ", sum1)
	fmt.Println("La suma de los numeros enviados es de: ", sum2)
}

func foo() {
	fmt.Println("Hola desde foo")
}

func bar(s string) {
	fmt.Println("Hola,", s)
}

func woo(st string) string {
	return fmt.Sprint("Hola desde woo,", st)
}

//funciones que retorna dos parametros, si no quiero retornar uno puedo poner _ para que no se envie
func saludar(n string, a string) (string, bool) {
	p := fmt.Sprint(n, " ", a, " dice hola ")
	q := true
	return p, q
}

//funcion que recibe parametros variables
func sumaNum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	suma := 0
	for _, v := range x {
		suma += v
	}
	return suma
}
