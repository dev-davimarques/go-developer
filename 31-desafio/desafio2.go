package main

import(
	"fmt"
)

func main(){
	for i := 1; i <=100; i++ {
		if i%3 == 0 {
			// fmt.Printf("%g é múltiplo de 3 [PIN]\n",i)
			fmt.Println(i," - é PIN [MÚLTIPLO DE 3]")
		} else if i%5==0 {
			// fmt.Printf("%g é múltiplo de 5 [PAN]\n",i)
			fmt.Println(i, " - é PAN [MÚLTIPLO DE 5]")
		}
	}
}