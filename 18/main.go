package main

import (
	"fmt"
	"sync"
)

type Count struct {
	count int
	mu    sync.Mutex
}

func (c *Count) Incriment() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Count) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	wg := &sync.WaitGroup{}
	counter := &Count{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Incriment()
			}
		}()
	}
	wg.Wait()
	fmt.Println("Result", counter.Value()) // Result 100000
}
