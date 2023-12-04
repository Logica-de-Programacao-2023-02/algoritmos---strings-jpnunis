package main

import (
	"fmt"
	"strings"
)

func findPatternIndices(text, pattern string) []int {
	var indices []int
	index := strings.Index(text, pattern)

	for index != -1 {
		indices = append(indices, index)
		index = strings.Index(text[index+1:], pattern)
		if index != -1 {
			index += (len(pattern) + 1)
		}
	}

	return indices
}

func main() {
	var text, pattern string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&text)
	fmt.Print("Digite um padrão: ")
	fmt.Scanln(&pattern)

	indices := findPatternIndices(text, pattern)

	if len(indices) > 0 {
		fmt.Printf("O padrão ocorre nos índices: %v\n", indices)
	} else {
		fmt.Println("O padrão não foi encontrado na string.")
	}
}
