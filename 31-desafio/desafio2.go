package main

import(
	"fmt"
)

func main(){
	for i := 1; i <=100; i++ {
		if i%3==0 {
			fmt.Println(i," - PIN")
		}
		if i%5==0 {
			fmt.Println(i," - PAN")
		}
		if i%3==0 && i%5==0 {
			fmt.Println(i," - PIN e PAN")
		}
	}
}