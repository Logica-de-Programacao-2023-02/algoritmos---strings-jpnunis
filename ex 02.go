package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	resultado := strings.ReplaceAll(input, " ", "")

	fmt.Println("Resultado:", resultado)
}
