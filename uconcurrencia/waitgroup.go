package main

import (
	"fmt"
	"runtime"
	"sync"
)

//declaramos una variable para waitgroup y poder accesar a sus metodos
var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(1)

	//pra hacer una go rutina separada, el proceso es muy rapido y no se logra imprimir
	go foo()
	bar()

	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}

	//cuando termina decrementa el valor de las gorutinas existentes
	//cuando al go rutina termina su trabajo decrementa el contador de wg.Add
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
