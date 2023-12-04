package main

import (
	"fmt"
	"strings"
)

func isIncreasingNumericSequence(input string) bool {
	for i := 0; i < len(input)-1; i++ {
		if input[i+1]-input[i] != 1 {
			return false
		}
	}

	return true
}

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	numericSequence := true
	for _, char := range input {
		if !strings.ContainsRune("0123456789", char) {
			numericSequence = false
			break
		}
	}

	if numericSequence && isIncreasingNumericSequence(input) {
		fmt.Println("É uma sequência numérica crescente!")
	} else {
		fmt.Println("Não é uma sequência numérica crescente.")
	}
}
