package main

import "fmt"

// AddFunc 定义一个函数类型
type AddFunc func(int64, int64) int64

// add 函数
func add(a, b int64) int64 {
	return a + b
}

// 使用闭包固定参数并加上偏移量
func fixedParameter(f AddFunc, offset int64) AddFunc {
	return func(a, b int64) int64 {
		return f(a, b) + offset
	}
}

// 或
//func fixedParameter(f func(int64, int64) int64, offset int64) func(int64, int64) int64 {
//	return func(a, b int64) int64 {
//		return f(a, b) + offset
//	}
//}

func main() {
	// 加上偏移量10
	fixedAdd := fixedParameter(add, 10)

	// 传入a,b
	result := fixedAdd(1, 1)
	fmt.Println(result)
}
