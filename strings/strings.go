package main

import "fmt"

func main() {
	//var s string
	//se pueden ocupar los unicode de escape y todo para los strings
	s := "Hello, world"
	//string es un arreglo de bytes acá se indica de que posicion hasta que posicion se va a imprimir
	s2 := s[0:5]
	s3 := s[7:11]
	//aca es desde el inicio del array hasta la posicion 5
	s4 := s[:5]
	//aca es desde la posicion 7 hata el final del array
	s5 := s[7:]
	fmt.Println(s, s2, s3, s4, s5)

	//arrays
	var vals [3]int
	vals[0] = 2
	vals[1] = 4
	vals[2] = 6
	fmt.Println(vals, vals[0], vals[1], vals[2])

	//no se pueden igular arrays de diferentes tamaños
	//var vals2 [4]int = vals
}
