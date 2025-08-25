package main

import (
	"fmt"
)

// 编写一个函数，给定两个通道 a 和 b ，返回一个相同类型的通道 c 。 a 或 b 接收到的每个元素都会被发送到 c ，并且一旦 a 和 b 都关闭， c 也会关闭。

var a = make(chan string)
var b = make(chan string)
var c = make(chan string)

func startA() {
	for i := 0; i < 5; i++ {
		a <- fmt.Sprintf("a%d", i)
	}
	close(a)
}

func startB() {
	for i := 0; i < 5; i++ {
		b <- fmt.Sprintf("b%d", i)
	}
	close(b)
}

func mergeAB() {
	aOpen, bOpen := true, true
	for aOpen || bOpen {
		select {
		case v, ok := <-a:
			if ok {
				c <- v
			} else {
				aOpen = false
			}
		case v, ok := <-b:
			if ok {
				c <- v
			} else {
				bOpen = false
			}
		default:
			fmt.Println(0)
		}
	}
	close(c)
}

func show() {
	for s := range c {
		fmt.Println(s)
	}
}
func main() {
	go startA()
	go startB()
	go mergeAB()
	show()
}

// 主要问题
//通道未关闭 ：startA 和 startB 函数发送完数据后没有关闭通道 a 和 b
//无限循环 ：mergeAB 函数中的 for 循环没有退出条件，会一直阻塞
//通道 c 未关闭 ：导致 show 函数中的 range 循环永远不会结束
//竞态条件 ：程序只等待1秒就退出，可能导致数据丢失
