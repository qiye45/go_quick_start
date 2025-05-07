package main

import (
	"fmt"
	"math"
)

// Point 结构体用于演示封装
// 封装了x,y坐标和相关操作方法
type Point struct {
	x, y float64 // 小写字段，包外不可见
}

// NewPoint 是Point的构造函数
func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

// Length 计算点到原点的距离
func (p Point) Length() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

// Distance 计算与另一点的距离
func (p Point) Distance(q Point) float64 {
	dx := p.x - q.x
	dy := p.y - q.y
	return math.Sqrt(dx*dx + dy*dy)
}

// Move 移动点的位置
func (p *Point) Move(dx, dy float64) {
	p.x += dx
	p.y += dy
}

// X 获取x坐标（getter）
func (p Point) X() float64 {
	return p.x
}

// Y 获取y坐标（getter）
func (p Point) Y() float64 {
	return p.y
}

// 自定义类型MyInt演示非结构体类型也可以有方法
type MyInt int

// Add 为MyInt类型添加方法
func (a MyInt) Add(b int) MyInt {
	return a + MyInt(b)
}

// Multiply 为MyInt类型添加方法
func (a MyInt) Multiply(b int) MyInt {
	return a * MyInt(b)
}

// DemoEncapsulation 演示Go语言中的封装
func DemoEncapsulation() {
	// 演示Point封装
	p := NewPoint(3, 4)
	fmt.Printf("点p(%.1f, %.1f)到原点的距离: %.1f\n", p.X(), p.Y(), p.Length())

	q := NewPoint(6, 8)
	fmt.Printf("点p到点q的距离: %.1f\n", p.Distance(q))

	p.Move(2, 3)
	fmt.Printf("移动后点p的坐标: (%.1f, %.1f)\n", p.X(), p.Y())

	// 演示自定义类型MyInt的方法
	var num MyInt = 10
	result1 := num.Add(5)
	result2 := num.Multiply(3)
	fmt.Printf("MyInt演示: %d + 5 = %d, %d * 3 = %d\n", num, result1, num, result2)
}
