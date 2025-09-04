package main

import (
	"sync"
	"time"
)

// GODEBUG=schedtrace=1000 go run ./godebug.go

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {
			a := 0

			for i := 0; i < 1e6; i++ {
				a += 1
			}

			wg.Done()
		}()
		time.Sleep(100 * time.Millisecond)
	}

	wg.Wait()
}
