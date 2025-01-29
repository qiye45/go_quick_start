package main

import "fmt"

func main() {
	// 使用make创建一个长度为3的字符串切片
	s := make([]string, 3)
	temp := make([]int, 3)
	fmt.Println("temp", temp)
	// 为切片的前三个元素赋值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	// 获取切片中索引为2的元素
	fmt.Println("get:", s[2]) // 输出：c

	// 获取切片的长度
	fmt.Println("len:", len(s)) // 输出：3

	// 使用append向切片追加一个元素"d"
	s = append(s, "d")

	// 使用append向切片一次性追加多个元素"e"和"f"
	s = append(s, "e", "f")

	// 打印完整的切片
	fmt.Println(s) // 输出：[a b c d e f]

	// 创建一个新切片c，长度与s相同
	c := make([]string, len(s))

	// 将s的内容复制到c
	copy(c, s)

	// 打印复制后的切片c
	fmt.Println(c) // 输出：[a b c d e f]

	// 切片操作：获取索引2到5(不包含5)的元素
	fmt.Println(s[2:5]) // 输出：[c d e]

	// 切片操作：获取开始到索引5(不包含5)的元素
	fmt.Println(s[:5]) // 输出：[a b c d e]

	// 切片操作：获取索引2到末尾的所有元素
	fmt.Println(s[2:]) // 输出：[c d e f]

	// 直接声明并初始化一个切片
	good := []string{"g", "o", "o", "d"}

	// 打印切片good
	fmt.Println(good) // 输出：[g o o d]
	fmt.Printf("%+v\n", good)
	//%v：默认格式
	//%+v：包含字段名的格式
	//%#v：Go语法格式
	//%T：类型信息
	//%p：指针地址
}
