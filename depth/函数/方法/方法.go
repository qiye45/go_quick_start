package main

type A struct {
	name string
}

func (a A) Name() string {
	a.name = "Hi " + a.name
	return a.name
}

func main() {
	a := A{name: "new world"}
	println(a.Name())
	println(A.Name(a))
}

func NameofA(a A) string {
	a.name = "Hi " + a.name
	return a.name
}
