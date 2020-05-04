package main

import "fmt"

func main() {
	//en go no es recomendable utilizar callbacks por que rompe el principio de programacion facil
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := suma(ii...)
	fmt.Println(s)
	//le pasamos la funcion suma, que la ejecutamos en el return
	s1 := sumaPar(suma, ii...)
	fmt.Println("El total de pares es:", s1)
}

func suma(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func sumaPar(f func(xi ...int) int, vi ...int) int {
	y := []int{}
	for _, v := range vi {
		if v%2 == 0 {
			y = append(y, v)
		}
	}
	return f(y...)
}
