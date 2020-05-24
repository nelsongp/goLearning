package main

import "fmt"

func main() {
	//los canales son para escribir codigo con currente de mejor forma
	//unbufferd chanel / canal sin buffer
	//solamente se puede utilzar cuando una go rutina envia y la otra recibe tiene que tener dos si o si
	ca := make(chan int)

	//go rotuna que envia el valor 42 al canal main
	go func() {
		//<- le digo que le envio 42 al canal
		ca <- 42
	}()

	fmt.Println(<-ca)
}
