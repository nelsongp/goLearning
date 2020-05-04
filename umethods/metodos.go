package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
}

type agenteSecreto struct {
	persona
	lpm bool
}

//a es el receptor del tipo que va a recibir y comunicador sera presentarse
func (a agenteSecreto) presentarse() {
	fmt.Println("Hola, soy", a.nombre, a.apellido)
}

func main() {
	as1 := agenteSecreto{
		persona: persona{
			nombre:   "nelson",
			apellido: "guardado",
		},
		lpm: true,
	}

	as2 := agenteSecreto{
		persona: persona{
			nombre:   "Juan",
			apellido: "Perez",
		},
		lpm: false,
	}
	fmt.Println(as1)
	as1.presentarse()

	//todos lso valores del mismo tipo tienen acceso al metodo especificado
	as2.presentarse()
}
