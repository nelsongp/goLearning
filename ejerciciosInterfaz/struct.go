package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
	edad     int
}

func main() {
	p := persona{
		nombre:   "nelson",
		apellido: "guardado",
		edad:     30,
	}

	p.presentarse()
}

//el metodo es una funcoin y reciben un receptor en este caso tipo persona
func (p persona) presentarse() {
	fmt.Println("Hola, soy", p.nombre, p.apellido, "y tengo", p.edad, "años")
}
