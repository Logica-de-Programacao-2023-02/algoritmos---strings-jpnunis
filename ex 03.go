package main

import (
	"fmt"
	"strings"
)

func main() {
	var input, oldChar, newChar string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	fmt.Print("Digite o caractere a ser substituído: ")
	fmt.Scanln(&oldChar)

	fmt.Print("Digite o novo caractere: ")
	fmt.Scanln(&newChar)

	resultado := strings.ReplaceAll(input, oldChar, newChar)

	fmt.Println("Resultado:", resultado)
}
