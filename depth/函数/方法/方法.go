package main

import (
	"fmt"
	"reflect"
)

type A struct {
	name string
}

func (a A) Name() string {
	a.name = "Hi " + a.name
	return a.name
}

func NameofA(a A) string {
	a.name = "Hi " + a.name
	return a.name
}

func main() {
	a := A{name: "new world"}
	println(a.Name())
	println(A.Name(a))

	t1 := reflect.TypeOf(A.Name)
	t2 := reflect.TypeOf(NameofA)

	fmt.Println(t1 == t2) // true
}
