/*
* Cálculo de Quadrados em Paralelo
*/

package main

import (
	"fmt"
	"time"
)

// Função que calcula o quadrado de um número
func calcularQuadrado(n int) {
	fmt.Printf("Calculando o quadrado de %d...\n", n)
	time.Sleep(1 * time.Second) // Simula uma operação demorada
	fmt.Printf("Resultado: %d\n", n*n)
}

func main() {
	// Iniciar goroutines para calcular o quadrado de vários números
	go calcularQuadrado(2)
	go calcularQuadrado(4)
	go calcularQuadrado(6)

	// Aguardar as goroutines terminarem
	time.Sleep(2 * time.Second) // Tempo suficiente para as goroutines concluírem
	fmt.Println("Fim do programa!")
}
