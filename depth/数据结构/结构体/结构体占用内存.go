package main

import (
	"fmt"
	"unsafe"
)

type s1 struct {
	a struct{}
}

type s2 struct {
	_ struct{}
}

type s3 struct {
	a struct{}
	b byte
}

type s4 struct {
	a struct{}
	b int64
}

type s5 struct {
	a byte     // 1字节
	b struct{} // 0字节（空结构体）
	c int64    // 8字节
}

// 1️⃣ Go 结构体的布局规则（和对齐相关）
// Go 的结构体字段在内存中排布时，会遵循几个原则：
//
// 每个字段的起始地址必须是它自身类型的对齐倍数
//
// byte 对齐值是 1
//
// int64 对齐值是 8
//
// 空结构体 struct{} 占 0 字节，对齐值是 1
//
// 字段之间会插入填充字节（padding） 以满足对齐要求
//
// 整个结构体的大小 必须是 对齐值最大的字段 的倍数
type s6 struct {
	a byte
	b struct{}
}

type s7 struct {
	a int64
	b struct{}
}

type s8 struct {
	a struct{}
	b struct{}
}

func main() {
	fmt.Println(unsafe.Sizeof(s1{})) // 0
	fmt.Println(unsafe.Sizeof(s2{})) // 0
	fmt.Println(unsafe.Sizeof(s3{})) // 1
	fmt.Println(unsafe.Sizeof(s4{})) // 8
	fmt.Println(unsafe.Sizeof(s5{})) // 16 内存对齐
	fmt.Println(unsafe.Sizeof(s6{})) // 2
	fmt.Println(unsafe.Sizeof(s7{})) // 16
	fmt.Println(unsafe.Sizeof(s8{})) // 0

	var a [10]int                 // 64位系统下是 8 字节
	fmt.Println(unsafe.Sizeof(a)) // 80

	var b [10]struct{}
	fmt.Println(unsafe.Sizeof(b)) // 0

	var c = make([]struct{}, 10)
	fmt.Println(unsafe.Sizeof(c)) // 24，即slice header的大小
}
