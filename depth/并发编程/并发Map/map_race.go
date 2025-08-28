package main

import (
	"fmt"
	"sync"
)

func map1(a map[int]int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		a[i] = i // 并发写入同一个map
	}
}

func main() {
	a := make(map[int]int)
	var wg sync.WaitGroup

	wg.Add(2)
	go map1(a, &wg)
	go map1(a, &wg)

	wg.Wait()
	fmt.Printf("map length: %d\n", len(a))
}
