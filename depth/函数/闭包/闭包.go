package main

func main() {
	fns := make([]func(), 0, 5)
	for i := 0; i < 5; i++ {
		fns = append(fns, func() {
			println(i)
		})
	}

	for _, fn := range fns { // 最后输出5个5，而不是0，1，2，3，4
		fn()
	}
}
