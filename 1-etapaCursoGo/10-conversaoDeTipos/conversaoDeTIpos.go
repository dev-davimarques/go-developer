package main

import (
	"fmt"
	"strconv"
)

func main(){
	// Variáveis
	var a int = 20
	var numInt int = 2024

	var c float64 = float64(a) // TIPO (VARIÁVEL)
	fmt.Println(c)
	c = 20.5
	fmt.Println(c)

	// Convertendo número em strings
	numeroEmString := strconv.Itoa(numInt)
	fmt.Printf("%q\n",numeroEmString)

}