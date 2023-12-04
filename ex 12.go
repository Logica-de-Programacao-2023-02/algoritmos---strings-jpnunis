package main

import (
	"fmt"
	"strings"
)

func isPalindrome(input string) bool {
	input = strings.ToLower(strings.ReplaceAll(input, " ", ""))
	length := len(input)

	for i := 0; i < length/2; i++ {
		if input[i] != input[length-1-i] {
			return false
		}
	}

	return true
}

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	if isPalindrome(input) {
		fmt.Println("É um palíndromo!")
	} else {
		fmt.Println("Não é um palíndromo.")
	}
}
