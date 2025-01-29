package main

import (
	"fmt"
	"strconv" // 导入字符串转换包
)

func main() {
	// ParseFloat 将字符串转换为浮点数
	// 第一个参数是要转换的字符串
	// 第二个参数是期望的位数（这里的64表示float64）
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f) // 输出: 1.234

	// ParseInt 将字符串转换为整数
	// 第一个参数是要转换的字符串
	// 第二个参数是进制（10表示十进制）
	// 第三个参数是结果的位数（64表示int64）
	n, _ := strconv.ParseInt("111", 10, 64)
	fmt.Println(n) // 输出: 111

	// 当进制参数为0时，ParseInt会根据字符串的前缀自动判断进制
	// 0x前缀表示十六进制
	n, _ = strconv.ParseInt("0x1000", 0, 64)
	fmt.Println(n) // 输出: 4096 (十六进制0x1000等于十进制4096)

	// Atoi 是 ParseInt(s, 10, 0) 的简写 "ASCII to integer"
	// 用于将字符串转换为十进制整数
	n2, _ := strconv.Atoi("123")
	fmt.Println(n2) // 输出: 123

	// 当转换失败时，Atoi 会返回错误
	// 这里尝试转换非数字字符串 "AAA"
	n2, err := strconv.Atoi("AAA")
	fmt.Println(n2, err) // 输出: 0 和错误信息
}
