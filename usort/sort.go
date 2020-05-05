package main

import (
	"fmt"
	"sort"
)

//type persona
type Persona struct {
	Nombre string
	Edad   int
}

type PorEdad []Persona

func (a PorEdad) Len() int           { return len(a) }
func (a PorEdad) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PorEdad) Less(i, j int) bool { return a[i].Edad < a[j].Edad }

func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Dr", "Moneluy"}

	sort.Ints(xi)
	sort.Strings(xs)

	fmt.Println(xi)
	fmt.Println(xs)

	p1 := Persona{"Eduar", 32}
	p2 := Persona{"Maria", 25}
	p3 := Persona{"Carolina", 56}
	p4 := Persona{"Adriana", 60}

	personas := []Persona{p1, p2, p3, p4}

	sort.Sort(PorEdad(personas))
	fmt.Println(personas)

}
