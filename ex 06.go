package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	palavras := strings.Fields(input)
	fmt.Printf("A string contém %d palavras.\n", len(palavras))
}
