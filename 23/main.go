package main

import "fmt"

func main() {
	slice := []int{2, 5, 1, 7, 6, 10, 22, 44, 1}
	var i int
	fmt.Scan(&i)
	fmt.Println(slice[i:])
	fmt.Println(slice[i+1:])
	copy(slice[i:], slice[i+1:])
	fmt.Println(slice)
	slice = slice[:len(slice)-1]
	fmt.Println(slice)
}
