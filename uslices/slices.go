package main

import "fmt"

func main() {
	//LITERAL COMPUESTO (COMPOSITE LITERAL)
	x := []int{1, 2, 3, 4, 5}

	//recorriendo un slice (for each)
	for i, v := range x {
		fmt.Printf("indice %v, valor %v \n", i, v)
	}

	//imprimir por rangos
	fmt.Println(x[2:4])

	//agregando elementos se le envía el array inicial y los valores a agregar
	x = append(x, 6, 7, 8)
	fmt.Println(x)

	//agregar un slice nuevo al slice existente
	y := []int{9, 10, 11}
	x = append(x, y...)
	fmt.Println(x)

	//quitar un elemento del slice
	x = append(x[:2], x[5:]...)
	fmt.Println(x)

	fmt.Println("--make array")
	//definicion de array con make se define el tipo del array, elementos a contener, capacidad total
	z := make([]int, 10, 12)
	fmt.Println(z)
	fmt.Println(len(z)) //cuenta numeros dentro del slice
	fmt.Println(cap(z)) //capacidad maxima del slice

	//esto no se peude hacer ya que se ha definido el slice con tamaño especifico aun qeu la capacidad sea mas grande
	//z[11] = 25
	//para agregar un nuevo elemento se debe de hacer lo siguiente y asi se incrementa el tamaña del slice sin desboradrlo
	z = append(z, 12)
	fmt.Println(z)

	//si agregamos un elemento fuera del tamaño definido lo que el hace es duplicar su capacidad ejemplo
	z = append(z, 13, 14, 15)
	fmt.Println(z)
	fmt.Println(len(z)) //cuenta numeros dentro del slice
	fmt.Println(cap(z)) //capacidad maxima del slice
	fmt.Println("--make array")

	//slice multidimencional
	vp := [][]int{x, z}
	fmt.Println(vp)

}
