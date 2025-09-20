package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func writer(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := range 1000 {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
				time.Sleep(2 * time.Second)
			}
		}
	}()
	return ch
}

func reader(ctx context.Context, ch <-chan int) {
	timeout := time.After(10 * time.Second)
	for {
		select {
		case <-ctx.Done():
			return
		case v, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println(v)
		case <-timeout:
			fmt.Println("Time exit")
			return
		}
	}
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	reader(ctx, writer(ctx))
}
