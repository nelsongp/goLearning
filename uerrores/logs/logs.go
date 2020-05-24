package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("sin-archivo.txt")
	if err != nil {
		//fmt.Println("ocurrion un problema", err)
		log.Println("un error ocurrio", err)
		//log hace que todo falle
		//log.Fatalln(err)
		//log panic termina la ejecucion de los llamados a f
		//log.Panicln(err)
	}
	defer f2.Close()
	fmt.Println("Revisa el archivo log.txt en el directorio")
}
