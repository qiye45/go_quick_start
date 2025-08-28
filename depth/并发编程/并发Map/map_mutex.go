package main

import (
	"fmt"
	"sync"
)

// 2. 互斥锁版本 (map_mutex.go)
//使用 sync.Mutex 保护普通map
//每次读写操作前加锁，操作后解锁
//保持原有map的使用方式

func map3(a map[int]int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mu.Lock()
		a[i] = i
		mu.Unlock()
	}
}

func main() {
	a := make(map[int]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)
	go map3(a, &mu, &wg)
	go map3(a, &mu, &wg)

	wg.Wait()
	fmt.Printf("map length: %d\n", len(a))
}
