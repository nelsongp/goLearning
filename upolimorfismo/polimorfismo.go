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

type manzana int

//a es el receptor del tipo que va a recibir y comunicador sera presentarse
//agente secreto recibe le metodo presentar, por lo tanto agente secreto implementa la interfaz humano

func (a agenteSecreto) presentarse() {
	fmt.Println("Hola, soy", a.nombre, a.apellido, "el agente secreto se presenta")
}

func (p persona) presentarse() {
	fmt.Println("Hola, soy", p.nombre, p.apellido, "la persona se presenta")
}

//Interfaz agente secreto implementa la interfaz humano
//la interfaz permite dar comportamientos
type humano interface {
	//metodo dentro de la interfaz
	presentarse()
}

func bar(h humano) {
	//Me devolvera el tipo de h en un dado caso agente secreto o persona
	//desde el llamado a la interfaz no puedo acceder a los metodos de los struc para hacerlo tengo que
	//hacer una accert de los tipos para poder acceder a sus metodos
	switch h.(type) {
	case persona:
		fmt.Println("Fui pasado a la funcion bar (persona)", h.(persona).nombre)
	case agenteSecreto:
		fmt.Println("Fui pasado a la funcion bar (agente secreto)", h.(agenteSecreto).nombre)
	}
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
	p := persona{
		nombre:   "Carito",
		apellido: "Guz",
	}

	fmt.Println(as1)
	as1.presentarse()

	//todos lso valores del mismo tipo tienen acceso al metodo especificado
	as2.presentarse()

	//aca se ve reflejado el polimorfismo ya que recibe dos tipos distintos
	bar(as1)
	bar(as2)
	bar(p)

	//conversion
	var x manzana = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
