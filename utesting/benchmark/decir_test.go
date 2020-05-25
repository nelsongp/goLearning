package main

import (
	"fmt"
	"testing"
)

func TestSaludar(t *testing.T) {
	l := "Bienvenido querido Eduar"
	s := saludar("Eduar")
	if s != l {
		t.Error("Expected", l, "Got", s)
	}

}

func ExampleSaludar() {
	fmt.Println(saludar("Eduar"))
	//Output:
	//"Bienvenido querido Eduar"
}

func BenchmarkSaludar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		saludar("Eduar")
	}
}
