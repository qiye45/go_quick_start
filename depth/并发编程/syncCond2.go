package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	flag bool
	cond *sync.Cond
	lk   sync.Mutex
)

func main() {
	cond = sync.NewCond(&lk)
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		for {
			time.Sleep(time.Second)
			cond.L.Lock()
			// 将flag 设置为true
			flag = true
			// 唤醒所有处于等待状态的协程
			cond.Broadcast()
			cond.L.Unlock()
		}
	}()

	for i := 0; i < 2; i++ {
		go func(i int) {
			defer wg.Done()
			for {
				time.Sleep(time.Millisecond * 500)
				cond.L.Lock()
				// 不满足条件，此时进入等待状态
				if !flag {
					cond.Wait()
				}
				// 被唤醒后，此时可能仍然不满足条件
				// 需要修改成
				//for !flag {
				//	cond.Wait()
				//}
				fmt.Printf("协程 %d flag = %t\n", i, flag)
				flag = false
				cond.L.Unlock()
			}
		}(i)
	}
	wg.Wait()
}
