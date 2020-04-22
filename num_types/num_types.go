package main

import "fmt"

func main() {
	var i int8 = 20
	var f float32 = 5.6
	//Parse values indicar el tipo al que se va a convertir
	fmt.Println(i + int8(f))
}
