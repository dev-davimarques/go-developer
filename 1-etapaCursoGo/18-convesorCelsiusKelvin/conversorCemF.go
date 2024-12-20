package main

import(
	"fmt"
)

func main(){
	fmt.Println(conversorF(25))
}

func conversorF(temp float64) float64 {
	result := (temp * 9/5)+32
	return result
}