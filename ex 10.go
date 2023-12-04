package main

import (
	"fmt"
	"strings"
)

func areAnagrams(str1, str2 string) bool {
	str1 = strings.ReplaceAll(strings.ToLower(str1), " ", "")
	str2 = strings.ReplaceAll(strings.ToLower(str2), " ", "")

	if len(str1) != len(str2) {
		return false
	}

	counts := make(map[rune]int)

	for _, char := range str1 {
		counts[char]++
	}

	for _, char := range str2 {
		counts[char]--
		if counts[char] < 0 {
			return false
		}
	}

	return true
}

func main() {
	var str1, str2 string
	fmt.Print("Digite a primeira string: ")
	fmt.Scanln(&str1)
	fmt.Print("Digite a segunda string: ")
	fmt.Scanln(&str2)

	if areAnagrams(str1, str2) {
		fmt.Println("Anagramas!")
	} else {
		fmt.Println("Não são anagramas.")
	}
}
