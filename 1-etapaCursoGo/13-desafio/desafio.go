package main

import(
	"fmt"
)

func main(){
	tempKelvin := 373.15
	tempCelsius := tempKelvin - 273.15
	fmt.Printf("A temperatura em Graus é: %g ºC\n",tempCelsius)
}