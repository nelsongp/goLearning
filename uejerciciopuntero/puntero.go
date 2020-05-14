package main

import "fmt"

type persona struct {
	nombre string
}

func (p *persona) hablar() {
	fmt.Println("Hola soy", p.nombre)
}

type humano interface {
	hablar()
}

func diAlgo(h humano) {
	h.hablar()
}

func main() {
	p1 := persona{
		nombre: "Nelson",
	}
	//No se puede mandar ya que se envia un tipo persona y tenemos que enviarle un puntero
	//diAlgo(p1)
	diAlgo(&p1)
	//como p1 tiene direccion tambien puedo invocarlo de esta forma
	p1.hablar()
}
