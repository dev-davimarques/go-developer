package main

import(
	"fmt"
)

func main(){
	// Exemplo string e int
	x := make(map[string]int)
	x["chave"] = 10
	fmt.Println(x["chave"])
	fmt.Println(x)

	// Exemplo int e int
	y := make(map[int]int)
	y[0] = 15
	fmt.Println(y[0])
	fmt.Println(y)

	// String to String
	elemento := make(map[string]string)
	elemento["H"] = "Hidrogênio"
	elemento["He"] = "Hélio"
	elemento["Li"] = "Lítio"
	fmt.Println(elemento["Li"])
	fmt.Println(elemento)
}