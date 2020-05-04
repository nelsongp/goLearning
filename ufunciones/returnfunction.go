package main

import "fmt"

func main() {
	//al dejar los dos parentisis indicamos que se ejecutaran las dos funciones
	fmt.Println(foofunc()())
}

func foofunc() func() int {
	return func() int {
		return 1492
	}
}
