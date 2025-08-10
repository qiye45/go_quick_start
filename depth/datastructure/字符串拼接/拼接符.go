package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	strSlices := []string{"h", "e", "l", "l", "o"}

	var all string
	for _, str := range strSlices {
		all += str
		sh := (*reflect.StringHeader)(unsafe.Pointer(&all))
		fmt.Printf("str地址：%p，all地址：%p，all底层字节数组地址=0x%x\n", &str, &all, sh.Data)
		// 获取底层字节数组地址 新版本获取方式
		data := uintptr(unsafe.Pointer(unsafe.StringData(all)))
		fmt.Printf("str地址：%p，all地址：%p，all底层字节数组地址=0x%x\n", &str, &all, data)

	}
	//	 从上面输出中可以发现str和all地址一直没有变，但是all的底层字节数组地址一直在变化，这说明拼接符 + 在拼接字符串时候，会创建许多临时字符串，临时字符串意味着内存分配，指向效率不会太高。
}
