package main

import (
	"fmt"
	"strings"
)

func main() {
	var input, oldChar, newChar string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	fmt.Print("Digite a letra a ser substituída: ")
	fmt.Scanln(&oldChar)

	fmt.Print("Digite a nova letra: ")
	fmt.Scanln(&newChar)

	resultado := strings.ReplaceAll(input, oldChar, newChar)

	fmt.Println("Resultado:", resultado)
}
