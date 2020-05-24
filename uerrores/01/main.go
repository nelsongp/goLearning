package main

import "fmt"

func main() {
	var respuesta1 string
	fmt.Print("Nombre:")
	_, err := fmt.Scan(&respuesta1)
	if err != nil {
		panic(err)
	}

	fmt.Println(respuesta1)
}
