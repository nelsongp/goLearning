package main

import (
	"fmt"
)

type persona struct {
	nombre   string
	apellido string
	sabores  []string
}

func main() {
	m := map[string]persona{}

	p1 := persona{
		nombre:   "Nelson",
		apellido: "Guardado",
		sabores: []string{"Arcoiris",
			"limon",
			"Fresa"},
	}

	p2 := persona{
		nombre:   "Marcela",
		apellido: "Morales",
		sabores: []string{"ChocoAlmendra",
			"Chocolate",
			"Limon"},
	}

	m["p1"] = p1
	m["p2"] = p2

	for _, v1 := range m {
		fmt.Println(v1.nombre, v1.apellido)
		for k, v := range v1.sabores {
			fmt.Println("\t", k, v)
		}
	}

}
