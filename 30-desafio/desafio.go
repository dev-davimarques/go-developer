package main

import(
	"fmt"
)

func main(){
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 {
			fmt.Println("Número divisível por 3: ",i)
		}
	}

	// for i := range 100{
	// 	if i % 3 == 0 {
	// 		fmt.Println("Número divisível por 3: ",i)
	// 	}
	// }
}