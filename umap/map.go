package main

import "fmt"

func main() {
	m := map[string]int{
		"Batman": 32,
		"Robin":  27,
	}
	fmt.Println(m)
	fmt.Println(m["Batman"])

	//si paso una llave que no existe me retorna el valor 0 enel mapa
	fmt.Println(m["Nelson"])

	//verificar si el valor de la llave existe en el mapa
	v, ok := m["Nelson"]
	fmt.Println(v)
	fmt.Println(ok)

	//Evaluamos si el valor existe en el mapa, de ser asi imprimimos la respuesta
	if v, ok := m["Robin"]; ok {
		fmt.Println("Imprimiendo desde el if", v)
	}

	//Ingrsamos al elemento nelson la edad 30
	m["Nelson"] = 30

	//recorrer mapa
	for k, v := range m {
		fmt.Println(k, v)
	}

	//borrar valor de un mapa
	delete(m, "Robin")
	fmt.Println(m)

	//si borramos una llaveque ya no esta el valor no retorna error para validar hacemos lo siguiente
	if v, ok := m["Batman"]; ok {
		fmt.Println("Se borró la llave con valor", v)
		delete(m, "Batman")
	}

	fmt.Println(m)
}
