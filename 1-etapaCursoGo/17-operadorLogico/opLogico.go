package main

import(
	"fmt"
)

func main(){
	x := 1
	if x == 2 || x == 3{
		fmt.Printf("Sim! %d é igual a 2 OU 3\n",x)
	} else {
		fmt.Printf("Não! %d não é igual a 2 OU 3\n",x)
	}

	y := 18
	if y%2 == 0 && y%3==0 {
		fmt.Printf("%d é divisível por 2 E por 3\n",y)
	} else {
		fmt.Println("%d não é divisível por 2 E por 3\n",y)
	}
}