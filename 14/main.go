package main

import "fmt"

func checkType(t interface{}) {
	switch v := t.(type) {
	case nil:
		fmt.Println("is nil")
	case int:
		fmt.Printf("type: %T, value: %v\n", v, v)
	case string:
		fmt.Printf("type: %T, value: %v\n", v, v)
	case bool:
		fmt.Printf("type: %T, value: %v\n", v, v)
	case chan interface{}:
		fmt.Printf("type: %T, value: %v\n", v, v)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	decimal := 111112
	str := "Maksim"
	boolean := false
	ch := make(chan interface{})

	checkType(decimal)
	checkType(str)
	checkType(boolean)
	checkType(ch)
}
