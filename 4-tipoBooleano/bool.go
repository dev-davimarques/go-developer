package main

import(
	"fmt"
)

func main(){
	// Operador lógico E = &&
	fmt.Println(true && true)
	fmt.Println(true && false)

	// Operador lógico OU = ||
	fmt.Println(true || true)
	fmt.Println(true || false)

	// Operador lógico Negação = !
	fmt.Println(!true)
}