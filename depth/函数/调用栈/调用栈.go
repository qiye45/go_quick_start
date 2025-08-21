package main

func sum(a, b int) int {
	sum := 0
	sum = a + b
	return sum
}

func main() {
	a := 3
	b := 5
	print(sum(a, b))
}
