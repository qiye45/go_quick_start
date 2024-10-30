package main

import (
	"fmt"
	"io"
	"os"
)

//这个通道的用途通常是管理和控制多个需要关闭的资源，例如文件、网络连接或数据库连接

func main() {
	processFiles()

}
func processFiles() {
	resources := make(chan io.Closer, 10) // 定义一个 io.Closer 类型的通道，缓冲区大小为 10

	// 创建多个文件资源并放入通道
	for i := 0; i < 3; i++ {
		file, _ := os.Open(fmt.Sprintf("file%d.txt", i))
		resources <- file
	}

	// 关闭所有资源
	closeAllResources(resources)
}

func closeAllResources(resources chan io.Closer) {
	for resource := range resources {
		resource.Close() // 逐个关闭资源
	}
}
