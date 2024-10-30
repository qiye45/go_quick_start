package main

import (
	"fmt"
	"time"
)

//假设你在家里等快递，有两个快递员可能会送来不同的包裹，但你不知道谁会先到。这时你可以同时等着，哪个快递员先敲门你就先收谁的包裹。这和 select 很像，它可以同时等待多个通道上的消息，哪个通道先收到数据，就执行哪个分支的代码。

func main() {
	// 创建两个通道，分别代表两个快递员
	package1 := make(chan string)
	package2 := make(chan string)

	// 启动两个 goroutine 模拟不同的送货时间
	go func() {
		time.Sleep(2 * time.Second) // 模拟送货延迟
		package1 <- "包裹1已送达！"       // 快递员1送达
	}()

	go func() {
		time.Sleep(1 * time.Second) // 模拟送货延迟
		package2 <- "包裹2已送达！"       // 快递员2送达
	}()

	// 使用 select 等待其中一个包裹先送达
	select {
	case msg := <-package1:
		fmt.Println(msg) // 如果 package1 先送达
	case msg := <-package2:
		fmt.Println(msg) // 如果 package2 先送达
	}
}
