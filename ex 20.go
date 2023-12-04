package main

import (
	"fmt"
	"unicode"
)

func isCamelCase(input string) (bool, int) {
	wordCount := 1

	for i, char := range input {
		if i == 0 && !unicode.IsUpper(char) {
			return false, 0
		}

		if unicode.IsUpper(char) {
			wordCount++
		}
	}

	return true, wordCount
}

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	isCamel, wordCount := isCamelCase(input)

	if isCamel {
		fmt.Printf("A string está em CamelCase e possui %d palavra(s).\n", wordCount)
	} else {
		fmt.Println("A string não está em CamelCase.")
	}
}
