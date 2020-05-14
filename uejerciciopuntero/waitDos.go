package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	//mutex funcoines lock y unlock que hace que las variables se bloquen y liberen a medida se usan en las go rutinas
	incremento := 0
	gs := 100
	wg.Add(gs)
	var m sync.Mutex
	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			v := incremento
			v++
			incremento = v
			fmt.Println(incremento)
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("El valor final del incremento es:", incremento)
}
