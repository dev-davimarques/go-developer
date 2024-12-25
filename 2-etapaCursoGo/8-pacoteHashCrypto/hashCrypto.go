package main

import(
	"fmt"
	// "hash/crc32"
	"crypto/sha1"
)

func main(){
	// Criar a hash
	// h := crc32.NewIEEE()
	h := sha1.New()

	// Escrever nossos dados no hash
	h.Write([]byte("Código com pacote hash"))

	// Calcular o hash
	v := h.Sum([]byte{})
	fmt.Println(v)
}