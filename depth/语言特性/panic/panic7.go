package main

func main() {
	defer doRecover()
	panic("it is panic") // recover
}

func doRecover() {
	defer recover()
}
