package main

import (
	"fmt"
	"strings"
)

func isDecreasingNumericSequence(input string) bool {
	for i := 0; i < len(input)-1; i++ {
		if input[i] != input[i+1]+1 {
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

	if numericSequence && isDecreasingNumericSequence(input) {
		fmt.Println("É uma sequência numérica decrescente!")
	} else {
		fmt.Println("Não é uma sequência numérica decrescente.")
	}
}
