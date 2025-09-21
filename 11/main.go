package main

import (
	"fmt"
	"math/rand/v2"
)

func intersect(a, b []int) []int {
	if len(b) < len(a) {
		a, b = b, a
	}
	seen := make(map[int]struct{}, len(a))
	for _, v := range a {
		seen[v] = struct{}{}
	}

	res := make([]int, 0, len(seen))
	for _, v := range b {
		if _, ok := seen[v]; ok {
			res = append(res, v)
			delete(seen, v)
		}
	}
	return res
}

func main() {
	a, b := make([]int, rand.IntN(100)), make([]int, rand.IntN(100))
	for i := range a {
		a[i] = rand.IntN(1000)
	}
	for i := range b {
		b[i] = rand.IntN(1000)
	}

	fmt.Println("Слайс a:", a)
	fmt.Println("Слайс b:", b)
	fmt.Println("Пересечение слайсов:", intersect(a, b))
}
