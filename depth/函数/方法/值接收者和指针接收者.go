package main

type A struct {
	name string
}

func (a A) GetName() string {
	return a.name
}

func (pa *A) SetName() string {
	pa.name = "Hi " + pa.name
	return pa.name
}

func main() {
	a := A{name: "new world"}
	pa := &a

	println(pa.GetName()) // 通过指针调用定义的值接收者方法
	println(a.SetName())  // 通过值调用定义的指针接收者方法
}
