package main

import (
	"fmt"
	"strings"
)

func replaceVowelsWithAsterisk(input string) string {
	vowels := "aeiouAEIOU"
	result := strings.Builder{}

	for _, char := range input {
		if strings.ContainsRune(vowels, char) {
			result.WriteRune('*')
		} else {
			result.WriteRune(char)
		}
	}

	return result.String()
}

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	result := replaceVowelsWithAsterisk(input)

	fmt.Println("Resultado após substituir as vogais por '*':", result)
}
