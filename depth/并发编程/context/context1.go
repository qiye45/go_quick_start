package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() { // 1
	fmt.Println("goroutine:", runtime.NumGoroutine())
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	go SlowOperation(ctx) // 2
	go func() {           // 3
		for {
			time.Sleep(300 * time.Millisecond)
			fmt.Println("goroutine:", runtime.NumGoroutine())
		}
	}()
	time.Sleep(2 * time.Second)

}

func SlowOperation(ctx context.Context) {
	done := make(chan int, 1)
	go func() { // 模拟慢操作 4
		dur := time.Duration(rand.Intn(2)+1) * time.Second // 1~2
		time.Sleep(dur)
		done <- 1
	}()

	select {
	case <-ctx.Done():
		fmt.Println("SlowOperation timeout:", ctx.Err())
	case <-done:
		fmt.Println("Complete work")
	}
}
