package main

import (
	"fmt"
	"math"
)

type circulo struct {
	radio float64
}

type cuadrado struct {
	longitud float64
}

func (c circulo) area() float64 {
	return math.Pi * c.radio * c.radio
}

func (c cuadrado) area() float64 {
	return c.longitud * c.longitud
}

//si no quiero acceder a ningun metodo diferente puedo nada mas hacer referencia a los metodos en la interfaz
type forma interface {
	area() float64
}

func info(f forma) {
	fmt.Println(f.area())
}

func main() {
	cua := cuadrado{
		longitud: 2,
	}

	cir := circulo{
		radio: 5,
	}

	info(cir)
	info(cua)

}
