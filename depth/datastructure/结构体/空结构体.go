package main

import (
	"fmt"
	"unsafe"
)

// Empty empty_struct.go
type Empty struct{}

// go:linkname zerobase runtime.zerobase
var zerobase uintptr // 使用go:linkname编译指令，将zerobase变量指向runtime.zerobase

// 内存逃逸分析
//  go run -gcflags="-m -l" 空结构体.go

func main() {
	a := Empty{}
	b := struct{}{}

	fmt.Println(unsafe.Sizeof(a) == 0) // true
	fmt.Println(unsafe.Sizeof(b) == 0) // true
	fmt.Printf("%p\n", &a)             // 0x590d00
	fmt.Printf("%p\n", &b)             // 0x590d00
	fmt.Printf("%p\n", &zerobase)      // 0x590d00

	c := new(Empty)
	d := new(Empty)
	_ = fmt.Sprint(c, d) // 目的是让变量c和d发生逃逸
	println(c)           // 0x590d00
	println(d)           // 0x590d00
	fmt.Println(c == d)  // true

	e := new(Empty)
	f := new(Empty)
	println(e)          // 0xc00008ef47
	println(f)          // 0xc00008ef47
	fmt.Println(e == f) // false
	// 可以看到变量 c 和 d 逃逸到堆上，它们打印出来的都是 0x591d00，且两者进行相等比较时候返回 true。而变量 e 和 f 打印出来的都是0xc00008ef47，但两者进行相等比较时候却返回false。这因为Go有意为之的，当空结构体变量未发生逃逸时候，指向该变量的指针是不等的，当空结构体变量发生逃逸之后，指向该变量是相等的。这也就是 Go官方语法指南 所说的：
}
