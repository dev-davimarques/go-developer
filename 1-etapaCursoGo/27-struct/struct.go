package main

import(
	"fmt"
)

func main(){
	// o que são estruturas? São coleção de campos
	// agrupar dados / formar registros

	type Biblioteca struct{
		Categoria string
		NomeLivro string
		Isbn int
	}

	livro1 := Biblioteca{Categoria: "terror", NomeLivro: "Frankenstein", Isbn: 1234567}

	fmt.Println(livro1)
	fmt.Println(livro1.NomeLivro)

	livros := []Biblioteca{}
	livros = append(livros, livro1)
	fmt.Println(livros)

	livro2 := Biblioteca{"Romance","O vento levou",1234556789}
	livros = append(livros, livro2)
	fmt.Println(livros)
}