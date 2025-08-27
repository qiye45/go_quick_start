package main

import (
	"fmt"
	"unsafe"
)

type Value struct {
	v interface{}
}

type ifaceWords struct {
	typ  unsafe.Pointer
	data unsafe.Pointer
}

func main() {
	val := Value{v: 123456}
	t := (*ifaceWords)(unsafe.Pointer(&val))
	dp := (*t).data            // dp是非安全指针类型变量
	fmt.Println(*((*int)(dp))) // 输出123456

	var val2 Value
	t = (*ifaceWords)(unsafe.Pointer(&val2))
	fmt.Println(t.typ) // 输出nil
}
