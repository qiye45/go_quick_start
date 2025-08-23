package main

func main() {
	ch := make(chan int, 1)
	select {
	case ch <- getVal(1):
		println("recv: ", <-ch)
	case ch <- getVal(2):
		println("recv: ", <-ch)
	case ch <- getVal(3):
		println("recv: ", <-ch)
	case ch <- getVal(4):
		println("recv: ", <-ch)
	}
}

func getVal(n int) int {
	println("getVal: ", n)
	return n
}
