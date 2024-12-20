package main

import(
	"fmt"
	"strings"
)

func main(){
	fmt.Println("Utilizando o pacote Strings - Função Contains")
	fmt.Println(strings.Contains("Computador", "dor"))
	fmt.Println(strings.Contains("Terra", "ceu"))

	meuNome := "Davi Marques"
	// DEIXAR EM CAIXA ALTA
	fmt.Println(strings.ToUpper(meuNome))
	// deixar em caixa baixa
	fmt.Println(strings.ToLower(meuNome))
	// Contar a quantidade de letras especificadas
	fmt.Println("Quantidade de a:",strings.Count(meuNome, "a"))
	// Quantidade total de caracteres (conta os espaços)
	fmt.Println(len(meuNome))
}