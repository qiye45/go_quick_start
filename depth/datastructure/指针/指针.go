package main

import "fmt"

type MyInt int
type PInt *int
type PMyInt *MyInt

func main() {
	var A int = 100
	var B *int = &A

	fmt.Println(A == *B)

	p1 := new(int)
	var p2 PInt = p1 // p2底层类型是*int
	p3 := new(MyInt)
	var p4 PMyInt = p3 // p4底层类型是*MyInt
	fmt.Println(p1, p2, p3, p4)
}
