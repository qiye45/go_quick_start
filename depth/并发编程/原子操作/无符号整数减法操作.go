package main

import (
	"sync/atomic"
)

func main() {
	var i uint64 = 100
	var j uint64 = 10
	var k = 5
	atomic.AddUint64(&i, -j)
	println(i)
	atomic.AddUint64(&i, -uint64(k))
	println(i)
	// 下面这种操作是不可以的，会发生恐慌：constant -5 overflows uint64
	// atomic.AddUint64(&i, -uint64(5))
}
