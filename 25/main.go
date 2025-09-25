package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func timeSleep(ns time.Duration) {
	<-time.After(ns)
}

func writer(ctx context.Context, wg *sync.WaitGroup) <-chan int {
	wg.Add(1)
	ch := make(chan int)
	go func() {
		defer wg.Done()
		defer close(ch)
		for i := 0; i < 100; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
				fmt.Printf("Writer sleeps for %d seconds\n", 5)
				timeSleep(5 * time.Second) // sleeping
			}
		}
	}()
	return ch
}
func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	wg := &sync.WaitGroup{}
	dataCh := writer(ctx, wg)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range dataCh {
			fmt.Printf("Reader: %d\n", v)
		}
	}()
	wg.Wait()
}
