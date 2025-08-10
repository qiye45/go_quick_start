package main

import (
	"fmt"
	"math/rand"
)

type randomOrder struct {
	count    uint32
	coprimes []uint32
}

type randomEnum struct {
	i     uint32
	count uint32
	pos   uint32
	inc   uint32
}

func (ord *randomOrder) reset(count uint32) {
	ord.count = count
	ord.coprimes = ord.coprimes[:0]
	for i := uint32(1); i <= count; i++ { // 初始化素数集合
		if gcd(i, count) == 1 {
			ord.coprimes = append(ord.coprimes, i)
		}
	}
}

func (ord *randomOrder) start(i uint32) randomEnum {
	return randomEnum{
		count: ord.count,
		pos:   i % ord.count,
		inc:   ord.coprimes[i%uint32(len(ord.coprimes))],
	}
}

func (enum *randomEnum) done() bool {
	return enum.i == enum.count
}

func (enum *randomEnum) next() {
	enum.i++
	enum.pos = (enum.pos + enum.inc) % enum.count
}

func (enum *randomEnum) position() uint32 {
	return enum.pos
}

func gcd(a, b uint32) uint32 { // 辗转相除法取最大公约数
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	arr := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	var order randomOrder
	order.reset(uint32(len(arr)))

	fmt.Println("====第一次随机遍历====")
	for enum := order.start(rand.Uint32()); !enum.done(); enum.next() {
		fmt.Println(arr[enum.position()])
	}

	fmt.Println("====第二次随机遍历====")
	for enum := order.start(rand.Uint32()); !enum.done(); enum.next() {
		fmt.Println(arr[enum.position()])
	}
}
