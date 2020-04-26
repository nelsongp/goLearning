package mapper

import (
	"github.com/rylans/getlang"
)

//Para instalar una libreria de tercero hay que poner go get "url de la libreria"

func Greet(s string) string {
	info := getlang.FromString(s)
	switch info.LanguageCode() {
	case "en":
		return "Hello!"
	case "de":
		return "Guten !!"
	case "fr":
		return "Bonjour"
	default:
		return "i dont know your language yet"
	}
}
