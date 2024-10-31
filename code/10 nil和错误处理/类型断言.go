package main

import (
	"fmt"
)

func main() {
	var i interface{}

	// 将一个整数赋值给接口
	i = 42

	// 使用类型断言提取整数值
	if v, ok := i.(int); ok {
		fmt.Println("Integer value:", v)
	} else {
		fmt.Println("Not an integer")
	}

	// 将一个字符串赋值给接口
	i = "Hello, Go!"

	// 再次使用类型断言提取字符串值
	if v, ok := i.(string); ok {
		fmt.Println("String value:", v)
	} else {
		fmt.Println("Not a string")
	}

	// 尝试将接口中的值断言为不同的类型
	if v, ok := i.(int); ok {
		fmt.Println("Integer value:", v)
	} else {
		fmt.Println("Not an integer")
	}
}
