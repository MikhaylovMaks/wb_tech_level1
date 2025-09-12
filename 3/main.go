package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"time"
)

func worker(id int, ch <-chan int) {
	for num := range ch {
		fmt.Printf("worker %d: %d\n", id, num)
	}
}

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Println("no integer")
		return
	}

	ch := make(chan int)
	for i := 1; i <= n; i++ {
		go worker(i, ch)
	}
	for {
		num, _ := rand.Int(rand.Reader, big.NewInt(1000))
		ch <- int(num.Int64())
		time.Sleep(500 * time.Millisecond)
	}
}
