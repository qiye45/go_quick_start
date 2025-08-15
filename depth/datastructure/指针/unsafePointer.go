package main

import (
	"fmt"
	"unsafe"
)

// Float64bits 通过非安全类型指针，将T1转换成T2
func Float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}

type MyInt int
type MyType struct {
	f1 uint8
	f2 int
	f3 uint64
}

func main() {
	// 将非安全类型指针转换成uintptr类型
	a := MyInt(100)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%x\n", uintptr(unsafe.Pointer(&a)))
	// 将非安全类型指针转换成uintptr类型，并进行算术运算
	s := MyType{f1: 10, f2: 20, f3: 30}
	f2UintPtr := uintptr(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Offsetof(s.f2)))
	fmt.Printf("%p\n", &s)
	fmt.Printf("%x %v\n", f2UintPtr, *(*int)(unsafe.Pointer(f2UintPtr))) // f2UintPtr = s地址 + 8

	arr := [3]int{1, 23, 45}
	fmt.Printf("%p\n", &arr)
	for i := 0; i < 3; i++ {
		addr := uintptr(unsafe.Pointer(uintptr(unsafe.Pointer(&arr[0])) + uintptr(i)*unsafe.Sizeof(arr[0])))
		fmt.Printf("%x %v\n", addr, *(*int)(unsafe.Pointer(addr)))
	}

}
