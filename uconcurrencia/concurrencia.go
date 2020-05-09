package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Numero Cpu's:", runtime.NumCPU())
	fmt.Println("Numero goRutintas:", runtime.NumGoroutine())
	contador := 0
	const gs = 100
	//contador de go rutinas
	var wg sync.WaitGroup
	wg.Add(gs)
	//bloquea a auna variable mientras una go rutina esta siendo ejecutada
	var mu sync.Mutex
	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			v := contador
			v++
			//yield permite que se ejecute otra go rutina en paralelo y ella espera a que termine
			runtime.Gosched()
			contador = v
			mu.Unlock()
			//cada vez que una go rutina finalice su ejecucion done reduce le contador de las pendientes
			wg.Done()
			//wait es el que ve e contador el vigila a que las rutinas terminen
			fmt.Println("Numero goRutintas:", runtime.NumGoroutine())
		}()
	}
	wg.Wait()
	fmt.Println("Cuenta:", contador)
}
