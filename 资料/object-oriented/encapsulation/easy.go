package main

import "fmt"

type P struct {
	A int
	b string
}

func (P) M1() {
	fmt.Println("P M1")
}

func (P) M2() {
	fmt.Println("P M2")
}

type Q struct {
	c [5]int
	D float64
}

func (Q) M2() {
	fmt.Println("Q M2")
}
func (Q) M3() {
	fmt.Println("Q M3")
}

func (Q) M4() {
	fmt.Println("Q M3")
}

type T struct {
	P
	Q
	E int
}

// M2 重写方法：在 T 中重写 M2 方法，明确调用哪个嵌入结构体的 M2。
func (t T) M2() {
	t.P.M2() // 或者 t.Q.M2()
}

func main() {
	var t T
	t.M1()
	//需要显式调用
	t.P.M2()
	t.Q.M2()
	// 或重写方法
	t.M2()
	t.M3()
	t.M4()
	println(t.A, t.D, t.E)
}
