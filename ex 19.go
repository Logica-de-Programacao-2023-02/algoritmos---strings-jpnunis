package main

import (
	"fmt"
	"strings"
)

func reverseString(input string) string {
	runes := []rune(input)
	length := len(runes)

	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}

	return string(runes)
}

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	result := reverseString(input)

	fmt.Println("String invertida:", result)
}
