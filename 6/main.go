package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// 1. Завершение по условию
func condition() {
	go func() {
		for i := range 5 {
			fmt.Println("condition:", i)
			time.Sleep(time.Second)
		}
		fmt.Println("condition gorutine done")
	}()
	time.Sleep(5 * time.Second)
}

// 2. Завершение через канал уведомления
func channel() {
	ch := make(chan struct{})
	go func() {
		for {
			select {
			case <-ch:
				fmt.Println("channel stop")
				return
			default:
				fmt.Println("channel working...")
				time.Sleep(time.Second)
			}
		}
	}()
	time.Sleep(3 * time.Second)
	ch <- struct{}{}
	time.Sleep(time.Second)
}

// 3. Завершение через context
func stopContext() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("context stop")
				return
			default:
				fmt.Println("context working")
				time.Sleep(time.Second)
			}
		}
	}()
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(time.Second)
}

// 4. Завершение через runtime.Goexit
func goexit() {
	go func() {
		fmt.Println("before goexit")
		runtime.Goexit()
		fmt.Println("its not print")
	}()
	time.Sleep(time.Second)
}

// 5. Завершение по time.After
func timeAfter() {
	go func() {
		timeout := time.After(3 * time.Second)
		for {
			select {
			case <-timeout:
				fmt.Println("time.After stop")
				return
			default:
				fmt.Println("time.After working...")
				time.Sleep(time.Second)
			}
		}
	}()
	time.Sleep(5 * time.Second)
}

// 6. Завершение по закрытию канала
func closeChanel() {
	ch := make(chan int)
	go func() {
		for v := range ch {
			fmt.Println("got:", v)
		}
		fmt.Println("close channel stop")
	}()
	for i := 0; i < 3; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	close(ch)
	time.Sleep(time.Second)
}

// 7. Завершение через time.Timer
func timer() {
	t := time.NewTimer(5 * time.Second)
	go func() {
		<-t.C
		fmt.Println("timer stop")
	}()
	time.Sleep(7 * time.Second)
}

func main() {
	fmt.Println("=== 1. stop by condition ===")
	condition()

	fmt.Println("\n=== 2. stop by channel ===")
	channel()

	fmt.Println("\n=== 3. stop by context ===")
	stopContext()

	fmt.Println("\n=== 4. stop by runtime.Goexit ===")
	goexit()

	fmt.Println("\n=== 5. stop by time.After ===")
	timeAfter()

	fmt.Println("\n=== 6. stop by closing channel ===")
	closeChanel()

	fmt.Println("\n=== 7. stop by time.Timer ===")
	timer()

	fmt.Println("\nAll examples finished.")
}
