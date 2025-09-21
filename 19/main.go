package main

import (
	"fmt"
)

func backWords(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	input := "главрыба 🐟"
	fmt.Println("Исходная строка:", input)
	fmt.Println("Перевёрнутая строка:", backWords(input))
}
