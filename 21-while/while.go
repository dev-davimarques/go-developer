package main

import(
	"fmt"
)

func main(){
	// o while não existe de forma oficial no go
	// porém podemos fazer um ajuste técnico
	i := 0
	for i < 10{
		fmt.Printf("%d é menor que 10\n",i)
		i++
	}

}