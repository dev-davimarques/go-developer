package main

import(
	"fmt"
)

func media(lista []float64) float64{
	total := 0.0
	for _, valor := range lista {
		total += valor
	}
	return total / float64(len(lista)) // interrompe a função
}

func main(){
	// Programa para calcular a média de uma turma
	lista := []float64{98,93,77,82,83} // lista de notas
	fmt.Println(media(lista))
}