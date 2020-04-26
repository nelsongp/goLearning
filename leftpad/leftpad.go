package leftpad

import (
	"strings"
	"unicode/utf8"
)

//la cariable solo podra ser vista en el paquete
var defaultChar = ' '

/*
la funcion format toma un string y un int y retorna el string
marcado hacia la izquierda con espacios y el tamaño del int.
si el string es mas largo que el especificado en el legth
el string original es retornado
*/
func format(s string, size int) string {
	return FormatRune(s, size, defaultChar)
}

/*en go existen dos niveles de visibildad pacquetes visibles y publicos
los tipos publicos pueden ser accesados desde otro lados cuando el paquete es importaddo
y los visibles solo pueden ser vistos dentro del propio paquete donde han sido declarado

para ser un paquete VISIBLES se tiene que declarar con inicial minuscula
para ser un paquete PUBLICO se debe de declarar con la inicial mayuscula
*/

/*
FormatRune toma un string, int y un rune y retorna el string
mas a la izquierda con el rune specifico y su tamaña, si el string
es mas largo que el espicifico el string original es eretornado
*/
func FormatRune(s string, size int, r rune) string {
	l := utf8.RuneCountInString(s)
	if l >= size {
		return s
	}

	out := strings.Repeat(string(r), size-1) + s
	return out
}
