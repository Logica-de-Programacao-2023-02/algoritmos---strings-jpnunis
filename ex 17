package main

import (
	"fmt"
)

func uniqueLetters(input string) string {
	charCount := make(map[rune]int)

	for _, char := range input {
		charCount[char]++
	}

	result := ""
	for _, char := range input {
		if charCount[char] == 1 {
			result += string(char)
		}
	}

	return result
}

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	result := uniqueLetters(input)

	fmt.Println("Letras únicas na string:", result)
}
