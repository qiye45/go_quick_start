package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func pow(a int) func(int) int {
	// 闭包，变量a作为闭包变量
	return func(b int) int {
		result := 1
		for i := 0; i < b; i++ {
			result *= a
		}
		return result
	}
}

func filter(a []int, fn func(int) bool) (result []int) {
	for _, v := range a {
		if fn(v) {
			result = append(result, v)
		}
	}

	return result
}

func generateInteger() func() int {
	ch := make(chan int)
	count := 0
	go func() {
		for {
			ch <- count
			count++
		}
	}()

	return func() int {
		return <-ch
	}
}
func generateN() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	// 函数赋值给一个变量
	fn := add
	fmt.Println(fn(1, 2)) // 3

	// 函数作为返回值
	powOfTwo := pow(2)         // 2的x次幂
	fmt.Println(powOfTwo(3))   // 8
	fmt.Println(powOfTwo(4))   // 16
	powOfThree := pow(3)       // 3的x次幂
	fmt.Println(powOfThree(3)) // 27
	fmt.Println(powOfThree(4)) // 81

	// 函数作为函数参数传递
	data := []int{1, 2, 3, 4, 5}
	// 传递奇数过滤器函子，过滤出奇数
	fmt.Println(filter(data, func(a int) bool {
		return a&1 == 1
	})) // 1, 3, 5
	// 过滤出偶数
	fmt.Println(filter(data, func(a int) bool {
		return a&1 == 0
	})) // 2, 4

	// 使用闭包函数构建一个生成器
	generate := generateInteger()
	fmt.Println(generate()) // 0
	fmt.Println(generate()) // 1
	fmt.Println(generate()) // 2

	generateSimple := generateN()
	fmt.Println(generateSimple())
	fmt.Println(generateSimple())
}
