/*
* Um exemplo real do uso de goroutines é quando você precisa executar várias tarefas simultaneamente * para melhorar o desempenho de um programa. Um caso típico seria fazer requisições HTTP a múltiplos * servidores de forma paralela para economizar tempo.
*/

package main

import (
	"fmt"
	"net/http"
	"time"
)

// Função para fazer uma requisição HTTP e retornar o status
func fetchStatus(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Erro ao acessar %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	// Enviar o status e o tempo de resposta para o canal
	ch <- fmt.Sprintf("%s -> %s (%.2fs)", url, resp.Status, time.Since(start).Seconds())
}

func main() {
	// Lista de URLs para consultar
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.stackoverflow.com",
		"https://www.golang.org",
	}

	// Criar um canal para receber os resultados
	ch := make(chan string)

	// Iniciar uma goroutine para cada URL
	for _, url := range urls {
		go fetchStatus(url, ch)
	}

	// Coletar os resultados
	fmt.Println("Resultados das requisições:")
	for range urls {
		fmt.Println(<-ch)
	}
}
