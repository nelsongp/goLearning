package main

import (
	"fmt"
	"sync"
)

func main() {
	par := make(chan int)
	impar := make(chan int)
	fanin := make(chan int)

	//enviar
	go enviar(par, impar)

	//recibir
	go recibir(par, impar, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("Finalizar")
}

//solo para recibir
func enviar(p, i chan<- int) {
	for j := 0; j < 100; j++ {
		if j%2 == 0 {
			p <- j
		} else {
			i <- j
		}
	}
	close(p)
	close(i)

}

//REcibir
func recibir(p, i <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range p {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range i {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}
