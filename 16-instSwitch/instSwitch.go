package main

import(
	"fmt"
)
 func main(){
	mes := 5
	switch{
	case mes == 1: 
		fmt.Println("Janeiro")
	case mes == 2: 
		fmt.Println("Fevereiro")
	case mes == 3:
		fmt.Println("Março")
	case mes == 4:
		fmt.Println("Abril")
	case mes == 5:
		fmt.Println("Maio")
	case mes == 6:
		fmt.Println("Junho")
	case mes == 7:
		fmt.Println("Julho")
	case mes == 8:
		fmt.Println("Agosto")
	case mes == 9:
		fmt.Println("Setembro")
	case mes == 10:
		fmt.Println("Outubro")
	case mes == 11:
		fmt.Println("Novembro")
	case mes == 12:
		fmt.Println("Dezembro")
	}
 }