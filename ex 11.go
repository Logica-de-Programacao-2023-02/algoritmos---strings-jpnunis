package main

import (
	"fmt"
	"strings"
)

func removeVowels(input string) string {
	vowels := "aeiouAEIOU"
	result := strings.Builder{}

	for _, char := range input {
		if !strings.ContainsRune(vowels, char) {
			result.WriteRune(char)
		}
	}

	return result.String()
}

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	result := removeVowels(input)

	fmt.Println("Resultado após remover as vogais:", result)
}
