package main

import (
	"fmt"
	"strings"
)

func unicElem(s string) bool {
	s = strings.ToLower(s)
	unicMap := make(map[rune]struct{})
	for _, elem := range s {
		if _, ok := unicMap[elem]; ok == true {
			fmt.Printf("Elem '%s' repeat \n", string(elem))
			return false
		}
		unicMap[elem] = struct{}{}
	}
	return true
}

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(unicElem(str))
}
