package main

import "fmt"

func main() {
	defer doRecover()
	panic("it is panic") // recover
}

func doRecover() {
	recover()
	fmt.Println("hello")
}
