package main

import(
	"fmt"
)

func main(){
	// var media float64
	var total float64 = 0
	var x [5]float64
	x[0] = 5
	x[1] = 5
	x[2] = 5
	x[3] = 5
	x[4] = 5

	for i := 0; i < len(x); i++ {
		total += x[i]
		fmt.Println(i," - ",x[i])
	}

	fmt.Println("Média: ",total / float64(len(x)))

}