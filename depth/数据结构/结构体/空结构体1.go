package main

import "fmt"

type MustKeydStruct struct {
	Name string
	Age  int
	_    struct{}
}

// 阻止unkeyed方式初始化结构体
func main() {
	persion := MustKeydStruct{Name: "hello", Age: 10}
	fmt.Println(persion)
	persion2 := MustKeydStruct{"hello", 10} //编译失败，提示： too few values in MustKeydStruct{...}
	fmt.Println(persion2)
}
