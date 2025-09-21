package main

import "fmt"

func uniqueStrings(a []string) map[string]struct{} {
	seen := make(map[string]struct{}, len(a))
	for _, v := range a {
		seen[v] = struct{}{}
	}
	return seen
}

func main() {
	a := []string{"cat", "cat", "dog", "cat", "tree"}
	res := uniqueStrings(a)
	for k, _ := range res {
		fmt.Print(k, " ")
	}
}
