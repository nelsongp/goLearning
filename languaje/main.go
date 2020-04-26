package main

import (
	"fmt"

	"./src/mapper"
)

func main() {
	fmt.Println(mapper.Greet("Howdy, what's mew?"))
	fmt.Println(mapper.Greet("Comment allez vous?"))
	fmt.Println(mapper.Greet("Wie geht es Ihnen"))
}
