package main

import (
	"fmt"
)

// createHugeString возвращает большую строку указанного размера.
func createHugeString(size int) string {
	return string(make([]byte, size))
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10) // 1 << 10 == 1024
	// Проблема была в том, что v[:100] лишь ссылается на первые 100 байт
	// большого буфера v. Это удерживает в памяти весь большой объект.
	// Решение — явно скопировать нужную часть в новый буфер:
	// Это позволит сборщику мусора освободить память,
	// занимаемую исходной большой строкой v после того, как она перестанет быть нужной.
	justString = string([]byte(v[:100]))
}

func main() {
	someFunc()
	fmt.Println(len(justString)) // выведет 100
}
