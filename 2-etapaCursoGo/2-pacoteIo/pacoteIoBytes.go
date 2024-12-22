package main

import (
	// "fmt"
	"log"
	"io"
	"os"
)

func main(){
	if _, err := io.WriteString(os.Stdout, "Bom dia"); err != nil{
		log.Fatal(err)
	}
}