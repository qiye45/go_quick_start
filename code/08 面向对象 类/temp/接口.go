package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

type martian struct {
	msg string
}

// 入参有实现talker函数
func (m martian) talk() string {
	return m.msg
}

// 接口作为入参
func upper(t talker) {
	u := strings.ToUpper(t.talk())
	fmt.Println(u)
}

func main() {
	m1 := martian{"msg"}
	fmt.Println(m1.talk())
	upper(martian{"original_msg"})
	upper(m1)
}
