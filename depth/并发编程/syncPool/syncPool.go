package main

import (
	"fmt"
	"sync"
)

type A struct {
	Name string
}

func (a *A) Reset() {
	a.Name = ""
}

var pool = sync.Pool{
	New: func() interface{} {
		return new(A)
	},
}

func main() {
	objA := pool.Get().(*A)
	objA.Reset() // 重置一下对象数据，防止脏数据
	defer pool.Put(objA)
	objA.Name = "test123"
	fmt.Println(objA)
}
