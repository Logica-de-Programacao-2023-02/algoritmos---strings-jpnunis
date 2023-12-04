package main

import (
	"fmt"
	"strings"
)

func containsOnlyNumbers(input string) bool {
	for _, char := range input {
		if !strings.ContainsRune("0123456789", char) {
			return false
		}
	}
	return true
}

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	if containsOnlyNumbers(input) {
		fmt.Println("A string contém apenas números!")
	} else {
		fmt.Println("A string não contém apenas números.")
	}
}
