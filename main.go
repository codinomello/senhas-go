package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	var tamanho int
	var maiúsculas, números, símbolos bool

	// solicitar entrada do usuário
	fmt.Print("Digite o tamanho da senha: ")
	fmt.Scanln(&tamanho)
	fmt.Print("Incluir letras maiúsculas? (true/false): ")
	fmt.Scanln(&maiúsculas)
	fmt.Print("Incluir números? (true/false): ")
	fmt.Scanln(&números)
	fmt.Print("Incluir símbolos? (true/false): ")
	fmt.Scanln(&símbolos)

	// gerar a senha
	senha, err := gerarSenha(tamanho, maiúsculas, números, símbolos)
	if err != nil {
		fmt.Println("erro ao gerar a senha:", err)
		return
	}

	// exibir a senha gerada
	fmt.Println("senha gerada:", senha)
}

// função para gerar a senha
func gerarSenha(tamanho int, maiúsculas, números, símbolos bool) (string, error) {
	// definindo os conjuntos de caracteres que podem ser usados
	var caracteres = "abcdefghijklmnopqrstuvwxyz"
	if maiúsculas {
		caracteres += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if números {
		caracteres += "0123456789"
	}
	if símbolos {
		caracteres += "!@#$%^&*()_-+=<>?/{}[]|"
	}

	// gerando a senha
	var senha string
	for i := 0; i < tamanho; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(caracteres))))
		if err != nil {
			return "", err
		}
		senha += string(caracteres[num.Int64()])
	}

	return senha, nil
}
