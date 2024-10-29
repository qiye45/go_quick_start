package main

import "fmt"

// 声明接口
type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

// user类型值的指针实现接口
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

// admin代表一个拥有权限的管理员用户
// admin类型嵌入了user类型
type admin struct {
	user
	level string
}

func main() {
	//创建admin用户
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com:",
		},
		level: "super",
	}
	// 新建时要创建user
	normal := admin{
		user{"1", "@"},
		"11",
	}

	//实现了接口notifier的内部类型user的方法被提升到了外部，所以此时admin类型也实现了接口
	sendNotification(&ad)
	sendNotification(&normal)
}

func sendNotification(n notifier) {
	n.notify()
}
