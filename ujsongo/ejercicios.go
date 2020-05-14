package main

import (
	"encoding/json"
	"fmt"
)

type usuario struct {
	Nombre string
	Edad   int
}

func main() {
	u1 := usuario{
		Nombre: "Eduar",
		Edad:   32,
	}
	u2 := usuario{
		Nombre: "Juan",
		Edad:   27,
	}

	usuarios := []usuario{u1, u2}
	fmt.Println(usuarios)

	bs, err := json.Marshal(usuarios)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

	s := `[{"Nombre":"Eduar","Edad":32},{"Nombre":"Juan","Edad":27}]`
	bis := []byte(s)

	var usu []usuario

	error := json.Unmarshal(bis, &usu)
	if error != nil {
		fmt.Println(error)
	}

	fmt.Println(usu)
}
