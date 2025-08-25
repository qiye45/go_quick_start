package main

import "sync"

var l sync.Mutex
var a string

func f() {
	a = "hello, world"
	l.Unlock() // 2
}

func main() {
	l.Lock() // 1
	go f()
	l.Lock() // 3
	print(a)
}
