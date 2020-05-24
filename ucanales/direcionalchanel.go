package main

import "fmt"

func main() {

	ca := make(chan int, 2)
	//ca := make(chan <- int, 2) //send
	//ca := make(<- chan int, 2)  //receive

	ca <- 42
	ca <- 43

	fmt.Println(<-ca)
	fmt.Println(<-ca)
	fmt.Println("------")
	fmt.Printf("%T\n", ca)
}
