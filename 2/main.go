package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	wg := &sync.WaitGroup{}
	wg.Add(len(arr))

	for _, elem := range arr {
		go func() {
			defer wg.Done()
			fmt.Println(elem * elem)
		}()
	}

	wg.Wait()
}
