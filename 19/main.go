package main

import (
	"fmt"
	"strings"
)

func backWords(str string) string {
	words := strings.Split(str, " ")
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}

func main() {
	input := "snow dog sun"
	fmt.Println("Исходная строка:", input)
	fmt.Println("Перевёрнутая строка:", backWords(input))
}
