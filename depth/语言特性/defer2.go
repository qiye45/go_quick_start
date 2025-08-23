package main

import "fmt"

func A() {
	fmt.Println("A")
	defer B()
	defer C()
}
func B() {
	fmt.Println("B")
}

func C() {
	fmt.Println("C")
	defer D()
}
func D() {
	fmt.Println("D")
}

func main() {
	A() // ACDB
}

// func A() {
//	deferproc(B) // 注册延迟函数B
//	deferproc(C) // 注册延迟函数C
//	deferreturn() // 开始执行延迟函数
//}
//
//func C() {
//	deferproc(D) // 注册延迟函数C
//	deferreturn() // 开始执行延迟函数
//}
