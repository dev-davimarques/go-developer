package main

import(
	"fmt"
)

func main(){
	// 1º Forma
	fmt.Println("1º Forma")
	for i := 1; i <= 10; i++ {
		fmt.Print(i," ")
	}
	
	// 2º Forma
	fmt.Println()
	fmt.Println("2º Forma")
	c := 1
	for c <= 10{
		fmt.Print(c," ")
		c = c + 1
	}
}