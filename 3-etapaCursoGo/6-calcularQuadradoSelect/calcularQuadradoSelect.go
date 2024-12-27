package main

import (
    "fmt"
    "time"
)

func calcularQuadrado(n int)int {
	return n * n
}

func main() {

    // c1 := make(chan string)
    // c2 := make(chan string)

	c3 := make(chan int) 
	c4 := make(chan int)

    go func() {
        time.Sleep(1 * time.Second)
        c3 <- calcularQuadrado(2)
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c4 <- calcularQuadrado(4)
    }()

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c3:
            fmt.Println("Quadrado do 1º número: ", msg1)
        case msg2 := <-c4:
            fmt.Println("Quadrado do 2º número: ", msg2)
        }
    }
}