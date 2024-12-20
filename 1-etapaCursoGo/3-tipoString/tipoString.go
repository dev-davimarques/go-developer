package main

import(
	"fmt"
)

func main(){
	nome := "Davi"
	fmt.Println("Seu nome é: ",nome)
	qtdNome := len(nome)
	fmt.Printf("Seu nome possui %d caracteres\n",qtdNome)
}