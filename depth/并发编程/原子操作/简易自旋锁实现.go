package main

import (
	"sync/atomic"
)

type spin int64

func (l *spin) lock() bool {
	for {
		if atomic.CompareAndSwapInt64((*int64)(l), 0, 1) {
			return true
		}
		continue
	}
}

func (l *spin) unlock() bool {
	for {
		if atomic.CompareAndSwapInt64((*int64)(l), 1, 0) {
			return true
		}
		continue
	}
}

func main() {
	s := new(spin)

	for i := 0; i < 5; i++ {
		s.lock()
		go func(i int) {
			println(i)
			s.unlock()
		}(i)
	}
	for {

	}
}
