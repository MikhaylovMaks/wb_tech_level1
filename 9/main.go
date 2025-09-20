package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func writer(ctx context.Context, arr []int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < len(arr); i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- arr[i]:
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return ch
}

func reader(ctx context.Context, ch <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-ch:
				if !ok {
					return
				}
				out <- v * 2
			}
		}
	}()
	return out
}

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	wg := &sync.WaitGroup{}
	arr := []int{1, 2, 3, 5, 6, 7, 12, 22, 444, 1}
	save := reader(ctx, writer(ctx, arr))
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range save {
			fmt.Println(v)
		}
	}()
	wg.Wait()
}
