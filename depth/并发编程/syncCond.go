package main

import (
	"fmt"
	"sync"
)

var (
	mu    sync.Mutex
	cond  = sync.NewCond(&mu)
	queue []int
)

func consume(id int) {
	for {
		mu.Lock()
		for len(queue) == 0 {
			cond.Wait()
		}
		x := queue[0]
		queue = queue[1:]
		mu.Unlock()
		fmt.Printf("Consumer %d got %d\n", id, x)
	}
}

func produce(x int) {
	mu.Lock()
	queue = append(queue, x)
	mu.Unlock()
	cond.Signal() // 唤醒一个消费者
}

func produceALL() {
	mu.Lock()
	for i := 0; i < 10; i++ {
		queue = append(queue, i)
	}
	mu.Unlock()
	// 唤醒所有消费者
	cond.Broadcast()
}

func main() {
	for i := 0; i < 3; i++ {
		go consume(i)
	}
	//for i := 0; i < 10; i++ {
	//	produce(i)
	//}
	produceALL()
	for {
	}
}
