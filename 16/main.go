package main

import (
	"fmt"
	"math/rand"
)

func quickSort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high) // находим индекс опорного элемента
		quickSort(arr, low, p-1)       // сортируем левую часть
		quickSort(arr, p+1, high)      // сортируем правую часть
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high] // выбираем опорный элемент (последний)
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i] // меняем местами
		}
	}
	// ставим pivot на своё место
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {
	arr := make([]int, rand.Intn(100))
	for i, _ := range arr {
		arr[i] = rand.Intn(1000)
	}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("Sorted:", arr)
}
