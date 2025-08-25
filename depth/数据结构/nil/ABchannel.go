package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 你的任务（如果你选择接受的话）是编写一个函数，给定两个通道 a 和 b ，返回一个相同类型的通道 c 。 a 或 b 接收到的每个元素都会被发送到 c ，并且一旦 a 和 b 都关闭， c 也会关闭。
// chan int
// 这是一个双向通道类型
// 可以进行发送和接收操作
// 既可以向通道发送数据 c <- value，也可以从通道接收数据 value := <-c
// <-chan int
// 这是一个只读通道类型
// 只能进行接收操作
// 只能从通道接收数据 value := <-c，不能向通道发送数据

func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}
func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok {
					fmt.Println("a is done")
					// 禁用 select 语句中的某个案例
					// 为了避免前面描述的繁忙循环，我们想禁用 select 语句的一部分。具体来说，当 a 闭合时，我们希望删除 case v, ok := <- a ，对于 b 也类似。但是该怎么做呢？
					// 正如我们在开头提到的，从 nil 通道接收数据会永远阻塞。所以，要禁用从通道接收数据的 case ，我们只需将该通道设置为 nil 即可！
					// 然后我们可以停止使用 adone 和 bdone ，而是检查 a 和 b 是否为 nil 。
					a = nil
					continue
				}
				c <- v
			case v, ok := <-b:
				if !ok {
					fmt.Println("b is done")
					b = nil
					continue
				}
				c <- v
			}
		}
	}()
	return c
}
func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	// range 语句会遍历通道中的所有值，直到通道关闭。但是谁关闭了通道呢？
	// 让我们在 Go 例程中添加一个 defer 语句，以确保通道最终被关闭。
	// 请注意， defer 语句位于新的 goroutine 中调用的匿名函数内部，而不是 mergeAB 内部。否则，一旦退出 mergeAB ， c 就会被关闭，向其发送值会导致 panic。
	for v := range c {
		fmt.Println(v)
	}

}
