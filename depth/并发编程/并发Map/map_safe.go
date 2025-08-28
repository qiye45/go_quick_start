package main

import (
	"fmt"
	"sync"
)

// 创建了两个并发安全的版本：
//1. sync.Map 版本 (map_safe.go)
//使用 sync.Map 类型，内置并发安全
//Store() 方法写入数据
//Range() 方法遍历计数

func map2(a *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		a.Store(i, i)
	}
}

func main() {
	var a sync.Map
	var wg sync.WaitGroup

	wg.Add(2)
	go map2(&a, &wg)
	go map2(&a, &wg)

	wg.Wait()

	count := 0
	a.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	fmt.Printf("map length: %d\n", count)
}
