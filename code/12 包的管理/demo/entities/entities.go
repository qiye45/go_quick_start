package entities

// User 在程序里定义一个用户类型
type user struct {
	Name  string
	Email string
}
type User0 struct {
	Name  string
	email string
}

// Admin 在程序里定义了管理员
type Admin struct {
	user   // 嵌入的类型是未公开的
	Rights int
}
