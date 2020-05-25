package main

import "fmt"

func main() {
	fmt.Println(saludar("Nelson"))
}

//saludar saluda a una persona
func saludar(s string) string {
	return fmt.Sprint("Bienvenido querido ", s)
}
