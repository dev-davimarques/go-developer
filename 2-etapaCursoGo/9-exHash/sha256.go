package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func gerarHashSenha(senha string) string {
	// Cria um novo hash SHA256
	hash := sha256.New()

	// Escreve a senha no hash
	hash.Write([]byte(senha))

	// Converte o hash para uma string hexadecimal
	return hex.EncodeToString(hash.Sum(nil))
}

func main() {
	// Senha original
	senha := "minhaSenhaSegura122!"

	// Gerar o hash da senha
	hashSenha := gerarHashSenha(senha)

	fmt.Println("Senha original:", senha)
	fmt.Println("Hash da senha:", hashSenha)

	// Simular a validação de uma senha
	senhaEntrada := "minhaSenhaSegura123!" // senha inserida pelo usuário
	hashEntrada := gerarHashSenha(senhaEntrada)

	if hashSenha == hashEntrada {
		fmt.Println("Senha válida!")
	} else {
		fmt.Println("Senha inválida!")
	}
}
