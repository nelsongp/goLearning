package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("Numero Cpu's:", runtime.NumCPU())
	fmt.Println("Numero goRutintas:", runtime.NumGoroutine())
	var contador int64
	const gs = 100
	//contador de go rutinas
	var wg sync.WaitGroup
	wg.Add(gs)
	//el atomic sirve para eliminar race condition

	for i := 0; i < 100; i++ {
		go func() {
			//recibe la direccion del valor de la variable que vamos aocupar
			atomic.AddInt64(&contador, 1)
			runtime.Gosched()
			fmt.Println("Contador", atomic.LoadInt64(&contador))
			wg.Done()
			//wait es el que ve e contador el vigila a que las rutinas terminen
			fmt.Println("Numero goRutintas:", runtime.NumGoroutine())
		}()
	}
	wg.Wait()
	fmt.Println("Cuenta:", contador)
}
