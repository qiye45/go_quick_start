package main

import "fmt"

// 显示切片的长度，容量信息
func dump(label string, slice []int) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}
func main() {
	temp := make([]int, 0)
	dump("temp", temp)
	temp = append(temp, 1, 2, 3)
	dump("temp", temp)
	// 类型转化
	num := 1
	fmt.Printf("%T\n", num)
	if _, ok := interface{}(num).(float64); ok {
		fmt.Println("ok", ok)
	}

	// 无效的类型断言: temp[0].(float64) (左侧为非接口类型 int)
	//if index, ok := temp[0].(float64); ok {
	//	fmt.Println(index)
	//}
}
