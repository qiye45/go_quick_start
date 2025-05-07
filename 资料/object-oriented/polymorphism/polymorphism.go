package polymorphism

import (
	"fmt"
	"math"
)

// Shape 是一个接口，定义形状的行为
type Shape interface {
	Area() float64
	Perimeter() float64
	Name() string
}

// Rectangle 矩形结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// Area 计算矩形面积
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter 计算矩形周长
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Name 返回形状名称
func (r Rectangle) Name() string {
	return "矩形"
}

// Circle 圆形结构体
type Circle struct {
	Radius float64
}

// Area 计算圆形面积
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter 计算圆形周长
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Name 返回形状名称
func (c Circle) Name() string {
	return "圆形"
}

// Triangle 三角形结构体
type Triangle struct {
	A, B, C float64 // 三边长度
}

// Area 计算三角形面积（海伦公式）
func (t Triangle) Area() float64 {
	s := (t.A + t.B + t.C) / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

// Perimeter 计算三角形周长
func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

// Name 返回形状名称
func (t Triangle) Name() string {
	return "三角形"
}

// PrintShapeInfo 打印形状信息
// 这个函数接收Shape接口，体现了多态
func PrintShapeInfo(s Shape) {
	fmt.Printf("形状: %s, 面积: %.2f, 周长: %.2f\n", s.Name(), s.Area(), s.Perimeter())
}

// DemoPolymorphism 演示Go语言中的多态
func DemoPolymorphism() {
	// 创建不同的形状
	rect := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 4}
	triangle := Triangle{A: 3, B: 4, C: 5}

	// 使用一个Shape类型切片存储不同的形状
	shapes := []Shape{rect, circle, triangle}

	// 使用多态调用方法
	for _, shape := range shapes {
		PrintShapeInfo(shape)
	}

	// 直接演示接口多态
	fmt.Println("\n直接使用接口变量演示多态:")
	var shape Shape = rect
	PrintShapeInfo(shape)

	shape = circle
	PrintShapeInfo(shape)

	shape = triangle
	PrintShapeInfo(shape)
}
