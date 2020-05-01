package main

import "fmt"

//Definicion del struct
type persona struct {
	nombre   string
	apellido string
	edad     int
}

type agenteSecreto struct {
	persona
	lpm bool
}

func main() {
	as1 := agenteSecreto{
		persona: persona{
			nombre:   "James",
			apellido: "Bond",
			edad:     32,
		},
		lpm: true,
	}

	p2 := persona{
		nombre:   "Marcela",
		apellido: "Morales",
		edad:     26,
	}

	//Struct anonimo se crea en momento de ejecucion
	p3 := struct {
		nombre   string
		apellido string
		edad     int
	}{
		nombre:   "Marcela",
		apellido: "Morales",
		edad:     26,
	}

	//con el struct el objeto padre puede acceder directamente a las propiedades del objeto hijo
	fmt.Println(as1)
	fmt.Println(as1.nombre, as1.apellido, as1.edad)
	fmt.Println(p2)
	fmt.Println(p3)
}
