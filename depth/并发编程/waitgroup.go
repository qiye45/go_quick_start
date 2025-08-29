package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // 等价于 Add(-1)
	fmt.Printf("Worker %d: start\n", id)
	time.Sleep(time.Second) // 模拟耗时任务
	fmt.Printf("Worker %d: done\n", id)
}

func main() {
	var wg sync.WaitGroup

	// Add(3)：设置计数器为 3，表示有 3 个工作要完成
	wg.Add(3)

	for i := 1; i <= 3; i++ {
		go worker(i, &wg) // 启动三个 goroutine
	}

	// Wait(): 阻塞当前 goroutine，直到计数器归零
	wg.Wait()
	fmt.Println("All workers done")
}
