package main

func main() {
	// 指针类型变量
	var p *int
	println(p == nil) // true
	a := 100
	p = &a
	println(p == nil) // false

	// 通道
	var ch chan int
	println(ch == nil) // true
	ch = make(chan int, 0)
	println(ch == nil) // false

	ch1 := make(chan int, 0)
	println(ch1 == nil) // false
	ch1 = nil
	println(ch1 == nil) // true

	// 切片
	var s []int           // 此时s是nil slice
	println(s == nil)     // true
	s = make([]int, 0, 0) // 此时s是empty slice
	println(s == nil)     // false

	// 映射
	var m map[int]int // 此时m是nil map
	println(m == nil) // true
	m = make(map[int]int, 0)
	println(m == nil) // false

	// 函数
	var fn func()
	println(fn == nil)
	fn = func() {
	}
	println(fn == nil)

	fun1()
}

func fun1() {
	//	nil 与 接口比较
	var p *int
	var i interface{}                   // (T=nil, V=nil)
	println(p == nil)                   // true
	println(i == nil)                   // true
	println("i == p", i == p)           // false
	var pi interface{} = interface{}(p) // (T=*int, V= nil)
	println(pi == nil)                  // false
	println(pi == i)                    // fasle
	println(p == i)                     // false。跟上面强制转换p一样。当变量和接口比较时候，会隐式将其转换成接口

	var a interface{} = nil                  // (T=nil, V=nil)
	println(a == nil)                        // true
	var a2 interface{} = (*interface{})(nil) // (T=*interface{}, V=nil)
	println(a2 == nil)                       // false
	var a3 interface{} = (interface{})(nil)  // (T=nil, V=nil)
	println(a3 == nil)                       // true
}
