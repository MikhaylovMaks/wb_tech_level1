package main

import "fmt"

func SetBit(n int64, i uint, value bool) int64 {
	mask := int64(1) << i
	if value {
		n = n | mask
	} else {
		n = n &^ mask
	}
	return n
}

func main() {
	var name int64 = 5
	fmt.Printf("Исходное число: %d (%04b)\n", name, name)

	name = SetBit(name, 0, false)
	fmt.Printf("После обнуления 1-го бита: %d (%04b)\n", name, name)

	// Устанавливаем 2-й бит в 1
	name = SetBit(name, 3, true)
	fmt.Printf("После установки 2-го бита: %d (%04b)\n", name, name)
}
