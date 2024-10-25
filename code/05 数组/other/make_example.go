package main

import "fmt"

func main() {
	temp := make([]int, 10)
	// make容量10，固定了
	for i := 0; i < 11; i++ {
		ints := append(temp, i)
		fmt.Println(ints)
	}
	fmt.Println(temp)

}
