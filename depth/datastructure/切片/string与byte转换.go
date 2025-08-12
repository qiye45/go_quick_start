package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func byteArrayToString(b []byte) string {
	return string(b)
}

func stringToByteArray(s string) []byte {
	return []byte(s)
}

func main() {
	//零拷贝公用内存
	//s := "Hello"
	// 这样生成的字符串来自堆上，不在只读段
	s := string([]byte{'H', 'e', 'l', 'l', 'o'})
	b := StringToBytes(s) // 零拷贝转换

	fmt.Printf("原字符串: %q\n", s)
	fmt.Printf("原字节切片: %v\n", string(b))

	// 修改字节切片
	b[0] = '1'
	fmt.Println("修改字节切片后：")
	fmt.Printf("字节切片: %v\n", string(b))
	fmt.Printf("字符串: %q\n", s) // 会发现字符串内容也变了
}

//go tool compile -N -l -S main.go

func StringToBytes(s string) (b []byte) {
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Len
	return b
}

func StringToBytes2(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}
