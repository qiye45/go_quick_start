package main

func main() {
	defer recover()
	panic("it is panic") // not recover
}
