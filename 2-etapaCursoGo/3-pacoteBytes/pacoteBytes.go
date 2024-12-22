package main

import(
	"fmt"
	"bytes"	
)

func main(){
	fmt.Printf("%s", bytes.ToTitle([]byte("maçã")))
	fmt.Printf("%s", bytes.Title([]byte("davi bezerra marques")))
}