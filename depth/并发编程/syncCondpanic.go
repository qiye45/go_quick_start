package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	count int
	cond  *sync.Cond
	mu    sync.Mutex
)

func main() {
	cond = sync.NewCond(&mu)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			time.Sleep(time.Second)
			count++
			cond.Broadcast()
		}
	}()

	go func() {
		defer wg.Done()
		for {
			time.Sleep(time.Millisecond * 500)
			//cond.L.Lock()
			for count%10 != 0 {
				cond.Wait()
			}
			fmt.Printf("count = %d", count)
			//cond.L.Unlock()
		}
	}()
	wg.Wait()
}
