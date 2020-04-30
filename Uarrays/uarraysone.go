package main

import "fmt"

func main() {
	//en go preferiblemente es de ocupar slices y no arreglos ya que tienen un mejor uso de memoria
	var x [5]int
	fmt.Println(x)
	fmt.Println(len(x))
}
