package main

import(
	"fmt"
)

func main(){
	// go aceita tanto dessa forma
	var nome string = "Davi"
	var idade int = 24
	var altura float64 = 1.85
	fmt.Println(nome,idade,altura)

	// mas também aceita dessa forma
	var nome2 = "Davi"
	var idade2 = 24
	var altura2 = 1.85
	fmt.Println(nome2,idade2,altura2)

	// a própria linguagem consegue entender o tipo criado
}