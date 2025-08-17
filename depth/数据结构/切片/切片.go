package main

import (
	"fmt"
	"unsafe"
)

func main() {
	type sliceHeader struct {
		array unsafe.Pointer // 底层数据数组的指针
		len   int            // 切片长度
		cap   int            // 切片容量
	}
	a := []byte{'h', 'e', 'l', 'l', 'o'}
	b := a[2:3]   // 创建一个从索引2到3的切片，容量默认延伸到底层数组末尾
	c := a[2:3:3] // 创建一个从索引2到3的切片，并指定容量为3（即不延伸到底层数组末尾）
	fmt.Println("切片a:", a)
	fmt.Println("切片b:", b)
	fmt.Println("切片c:", c)
	ptrA := (*sliceHeader)(unsafe.Pointer(&a))
	ptrB := (*sliceHeader)(unsafe.Pointer(&b))
	ptrC := (*sliceHeader)(unsafe.Pointer(&c))

	fmt.Printf("切片%s: 底层数组地址=0x%x, 长度=%d, 容量=%d\n", "a", ptrA.array, ptrA.len, ptrA.cap)
	fmt.Printf("切片%s: 底层数组地址=0x%x, 长度=%d, 容量=%d\n", "b", ptrB.array, ptrB.len, ptrB.cap)
	fmt.Printf("切片%s: 底层数组地址=0x%x, 长度=%d, 容量=%d\n", "c", ptrC.array, ptrC.len, ptrC.cap)
}
