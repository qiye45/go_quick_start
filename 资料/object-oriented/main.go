package main

import (
	"fmt"
	"object-oriented/composition"
	"object-oriented/encapsulation"
	"object-oriented/polymorphism"
)

func main() {
	fmt.Println("=== Go \"无对象\" OO编程演示 ===")

	fmt.Println("\n=== 1. 封装示例 ===")
	encapsulation.DemoEncapsulation()

	fmt.Println("\n=== 2. 组合(\"继承\")示例 ===")
	composition.DemoComposition()

	fmt.Println("\n=== 3. 多态示例 ===")
	polymorphism.DemoPolymorphism()
}
