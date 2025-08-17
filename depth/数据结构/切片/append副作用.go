package main

import "fmt"

func main() {
	slice1 := []byte{'h', 'e', 'l', 'l', 'o'}
	slice2 := slice1[2:3]
	slice3 := slice1[2:3:3]
	fmt.Printf("slice1: %p\n", slice1)
	fmt.Printf("slice2: %p\n", slice2)
	fmt.Printf("slice3: %p\n", slice3)
	slice2 = append(slice2, 'g')
	slice3 = append(slice3, 'g') // append会复制一份底层数组，并返回新的切片
	fmt.Printf("slice3: %p  %+v\n", slice3, slice3)
	fmt.Println("len:", len(slice1), len(slice2), len(slice3))
	fmt.Println("cap:", cap(slice1), cap(slice2), cap(slice3))
	fmt.Println(string(slice1)) // 输出helge，slice1的值也变了。
	fmt.Println(string(slice2)) // lg
	fmt.Println(string(slice3)) // lg
	slice4 := make([]byte, 1)   // 创建一个长度为1的切片
	copy(slice4, slice1[2:3])
	slice4 = append(slice4, 'g')
	fmt.Println(string(slice4)) // 输出lg
	//slice5 := make([]byte, 1, 3) // len=1, cap=3

}
