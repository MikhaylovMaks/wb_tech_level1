package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// signal.NotifyContext оборачивает context.Background() и автоматически вызывает cancel(), если придёт Ctrl+C (SIGINT) или SIGTERM.

// В writer и reader мы проверяем <-ctx.Done(). Как только сигнал получен, контекст закрывается → горутины завершаются.

// sync.WaitGroup гарантирует, что main дождётся завершения всех ридеров.
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

func reader(ctx context.Context, ch <-chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-ch:
				if !ok {
					return
				}
				fmt.Println(v)
			}
		}
	}()
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	wg := &sync.WaitGroup{}
	ch := writer(ctx)
	reader(ctx, ch, wg)
	wg.Wait()
}
