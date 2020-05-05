package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Nombre   string `json:"Nombre"`
	Apellido string `json:"Apellido"`
	Edad     int    `json:"Edad"`
}

func main() {
	//decoder cuando no se quiere guardar informacion en varibales
	//marshal y unmarshal para cuando se quiere guardar algo
	//convertir json a go
	s := `[{"Nombre":"James","Apellido":"Bond","Edad":32},{"Nombre":"Miss","Apellido":"MOney","Edad":27}]`
	bs := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	//personas := []persona{}
	var personas []persona

	//En el valor se le envia el puntro el espacio en memoria del valor parseado &personas2
	err := json.Unmarshal(bs, &personas)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Toda la data ", personas)

	for i, v := range personas {
		fmt.Println("\nPERSONA NUMERO", i)
		fmt.Println(v.Nombre, v.Apellido, v.Edad)
	}
}
