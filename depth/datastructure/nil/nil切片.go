package main

func main() {
	var s []int
	//s[0] = 1           // panic: runtime error: index out of range [0] with length 0
	//println(s[0])      // panic: runtime error: index out of range [0] with length 0
	s = append(s, 100) // ok

}
