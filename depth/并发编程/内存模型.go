package main

import "time"

var a string

func hello() {
	go func() { a = "hello" }() // 1
	print(a)                    // 2
}

func main() {
	var a, b int

	// goroutine A
	go func() {
		a = 1
		b = 2
	}()

	// goroutine B
	go func() {
		println("a=", a)
		if b == 2 {
			println("b=2")
		} else {
			println("b!=2")
		}
	}()

	hello()
	time.Sleep(time.Second)
}

//当执行goroutine B打印变量a时并不一定打印出来1，有可能打印出来的是0。这是因为goroutine A中可能存在指令重排，先将b变量赋值2，若这时候接着执行goroutine B那么就会打印出来0
// a= 0
// b!=2
