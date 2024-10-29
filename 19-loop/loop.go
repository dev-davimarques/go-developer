package main

import(
	"fmt"
)

func main(){
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }
	for i := range 10{
		fmt.Println("range", i)
	}
}