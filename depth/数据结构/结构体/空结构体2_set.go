package main

import (
	"fmt"
)

type Set struct {
	items map[interface{}]emptyItem
}

type emptyItem struct{}

var itemExists = emptyItem{}

func NewSet() *Set {
	set := &Set{items: make(map[interface{}]emptyItem)}
	return set
}

// Add 添加元素到集合
func (set *Set) Add(item interface{}) {
	set.items[item] = itemExists
}

// Remove 从集合中删除元素
func (set *Set) Remove(item interface{}) {
	delete(set.items, item)

}

// Contains 判断元素是否存在集合中
func (set *Set) Contains(item interface{}) bool {
	_, contains := set.items[item]
	return contains
}

// Size 返回集合大小
func (set *Set) Size() int {
	return len(set.items)
}

func main() {
	set := NewSet()
	set.Add("hello")
	set.Add("world")
	set.Add(1)
	set.Add(nil)
	fmt.Println(set.Contains("hello"))
	fmt.Println(set.Contains("Hello"))
	fmt.Println(set.Size())
}
