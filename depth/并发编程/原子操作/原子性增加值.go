package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var count int32
	var unsafeCount int32
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&count, 1) // 原子性增加值
			wg.Done()
		}()
		go func() {
			fmt.Println(atomic.LoadInt32(&count)) // 原子性加载
		}()
		go func() {
			unsafeCount += 1
		}()
	}
	wg.Wait()
	fmt.Println("count: ", count)             // 100
	fmt.Println("unsafeCount: ", unsafeCount) // 99
}
