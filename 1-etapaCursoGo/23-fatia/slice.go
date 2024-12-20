package main

import(
	"fmt"
)

func main(){
	// var x []float64 -> Forma de declarar um array
	// fatia := make([]float64 ,4) -> Forma de criar um slice
	fatia := make([]string,2)
	fatia[0] = "Davi"
	fatia[1] = "Marques"
	fatia = append(fatia, "Teste")
	// fatia[2] = "20"
	// fatia[3] = "2025"
	// fatia[4] = "Teste"
	fmt.Println(len(fatia))
	fmt.Println(fatia)
	fmt.Println(fatia[0])
}