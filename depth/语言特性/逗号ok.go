package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	// 启动三个写数据的 goroutine
	go func() {
		for i := 1; i <= 3; i++ {
			time.Sleep(200 * time.Millisecond)
			ch1 <- i * 10
			fmt.Printf("[Writer1] 写入: %d\n", i*10)
		}
		close(ch1)
	}()

	go func() {
		for i := 1; i <= 3; i++ {
			time.Sleep(300 * time.Millisecond)
			ch2 <- i * 100
			fmt.Printf("[Writer2] 写入: %d\n", i*100)
		}
		close(ch2)
	}()

	go func() {
		for i := 1; i <= 2; i++ {
			time.Sleep(500 * time.Millisecond)
			ch3 <- i * 1000
			fmt.Printf("[Writer3] 写入: %d\n", i*1000)
		}
		close(ch3)
	}()

	// 主循环：同时监听三个 channel
	doneCount := 0
	for doneCount < 3 { // 等待三个 channel 都关闭
		select {
		case v, ok := <-ch1:
			if ok {
				fmt.Printf("[Main] 从 ch1 收到: %d\n", v)
			} else {
				ch1 = nil // 避免 select 再次监听已关闭的 ch
				doneCount++
			}
		case v, ok := <-ch2:
			if ok {
				fmt.Printf("[Main] 从 ch2 收到: %d\n", v)
			} else {
				ch2 = nil
				doneCount++
			}
		case v, ok := <-ch3:
			if ok {
				fmt.Printf("[Main] 从 ch3 收到: %d\n", v)
			} else {
				ch3 = nil
				doneCount++
			}
		default:
			// 如果所有 channel 此刻都没有数据，就不会阻塞
			time.Sleep(100 * time.Millisecond)
		}
	}

	fmt.Println("所有 channel 已关闭，程序结束")
}
