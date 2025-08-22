package main

type A struct {
	name string
}

func (a A) GetName() string {
	return a.name
}
func GetFunc() func() string {
	a := A{name: "new world"}
	return func() string {
		return A.GetName(a)
	}
}

func main() {
	a := A{name: "new world"}

	f1 := A.GetName // 方法表达式
	f1(a)

	f2 := a.GetName // 方法变量
	f2()

	f2 = GetFunc()
}
