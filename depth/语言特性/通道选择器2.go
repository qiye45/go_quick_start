package main

import "time"

func main() {
	ch := make(chan int, 1)
	go func() {
		time.Sleep(time.Second)
		ch <- 100
	}()

	for {
		select {
		case i := <-ch:
			println("case1 recv: ", i)
			return
		case i := <-ch:
			println("case2 recv: ", i)
			return
		default:
			println("default case")
			time.Sleep(time.Millisecond * 500)
		}
	}
}
