package main

//en este caso si se ocupa la clase math y la math/ rand tienen que importarse los dos siempre, ya que uno este dentro de otro no quiere
//decir que lo va a importar
import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	//se inicializa la variable rand con seed
	rand.Seed(time.Now().UnixNano())
	//en a y b se inicializan con numeros random mayors a 0 e iguales a 10
	a := rand.Intn(10)
	b := rand.Intn(10)
	//para trabajara con la clase math hay que pasara los valores a int64 para trabajar con sus funciones
	c := int(math.Max(float64(a), float64(b)))
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c, "is bigger")
}
