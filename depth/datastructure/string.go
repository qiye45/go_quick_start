package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := "hello"
	b := a
	fmt.Printf("a变量地址：%p\n", &a)
	fmt.Printf("b变量地址：%p\n", &b)
	a = "1"
	fmt.Println(a, b)
	print("断点打在这里")
	fmt.Println(unsafe.Sizeof([3]string{}))
}
