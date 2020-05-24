package main

import "fmt"

func main() {
	//los canales son para escribir codigo con currente de mejor forma
	//buffered chanel / canal con buffer
	//se le pasa como segundo elemento al make la cantidad de valores que enviare en el buffer
	ca := make(chan int, 2)

	ca <- 42
	ca <- 43

	fmt.Println(<-ca)
	fmt.Println(<-ca)
}
