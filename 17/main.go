package main

import (
	"fmt"
)

func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid // нашли элемент
		} else if arr[mid] < target {
			low = mid + 1 // ищем справа
		} else {
			high = mid - 1 // ищем слева
		}
	}

	return -1 // не найден
}

func main() {
	arr := []int{3, 10, 18, 29, 33, 42, 55, 71}
	fmt.Println("Array:", arr)

	fmt.Println("Search 29:", binarySearch(arr, 29))
	fmt.Println("Search 100:", binarySearch(arr, 100))
}
