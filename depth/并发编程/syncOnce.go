package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var config map[string]string

// 模拟加载配置
func loadConfig() {
	fmt.Println("加载配置文件...")
	config = map[string]string{
		"host": "127.0.0.1",
		"port": "8080",
	}
}

func getConfig() map[string]string {
	// 只会执行一次 loadConfig()
	once.Do(loadConfig)
	return config
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	// 多个 goroutine 并发调用
	for i := 1; i <= 3; i++ {
		go func(id int) {
			defer wg.Done()
			cfg := getConfig()
			fmt.Printf("Goroutine %d 读取配置: %v\n", id, cfg)
		}(i)
	}

	wg.Wait()
}
