package main

import "fmt"

// defer函数的传入参数在定义时就已经明确
func fun1() {
	i := 1
	defer fmt.Println(i)
	i++
}

// defer函数是按照后进先出的顺序执行
func fun2() {
	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}
}

// defer函数可以读取和修改函数的命名返回值
func fun3() (i int) {
	defer func() {
		i++
	}()
	return 100
}
func main() {
	fun1()
	fun2()
	fmt.Println(fun3())

}
