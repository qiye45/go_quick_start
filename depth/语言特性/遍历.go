package main

import (
	"fmt"
	"unicode/utf8"
)

type User struct {
	Name string
	Age  int
}

func main() {
	var a [3]int
	for i, v := range a {
		fmt.Println(i, v)
	}

	for i, v := range &a {
		fmt.Println(i, v)
	}
	//	遍历切片
	b := make([]int, 3)
	for i, v := range b {
		fmt.Println(i, v)
	}

	b = nil
	for i, v := range b {
		fmt.Println(i, v)
	}
	// 当遍历切片时候，可以边遍历边append操作，这并不会造成死循环
	c := make([]int, 3)
	for _, num := range c {
		c = append(c, num)
	}
	fmt.Printf("%v\n", c)

	// for-range切片时候，返回的是值拷贝
	users := []User{
		{
			Name: "a1",
			Age:  100,
		},
		{
			Name: "a2",
			Age:  101,
		},
		{
			Name: "a2",
			Age:  102,
		},
	}

	fmt.Println("before: ", users)
	for _, v := range users {
		v.Age = v.Age + 10 // 想给users中所有用户年龄增加10岁
	}
	fmt.Println("after:  ", users)
	//	使用索引修改原数组
	for i := range users {
		users[i].Age = users[i].Age + 10
	}
	fmt.Println("after:  ", users)

	//	rune类型，由于遍历字符串时候，返回的是码点，所以索引并不总是依次增加1的
	var str = "hello，你好"
	var buf [100]byte
	for i, v := range str {
		vl := utf8.RuneLen(v)
		si := i + vl
		copy(buf[:], str[i:si])
		fmt.Printf("索引%2d: %q，\t 码点: %#6x，\t 码点转换成字节: %#v\n", i, v, v, buf[:vl])
	}

	//	遍历映射，不会顺序遍历
	m := map[int]int{
		1: 10,
		2: 20,
		3: 30,
		4: 40,
		5: 50,
	}

	for i, v := range m {
		fmt.Println(i, v)
	}

	m = nil
	for i, v := range m {
		fmt.Println(i, v)
	}
}
