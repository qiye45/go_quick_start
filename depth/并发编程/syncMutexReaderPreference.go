package main

import (
	"fmt"
	"sync"
	"time"
)

// 你实现的 rwlock
type rwLock struct {
	readerCnt  int
	readerLock sync.Mutex
	writerLock sync.Mutex
}

func NewRWLock() *rwLock {
	return &rwLock{}
}

func (l *rwLock) RLock() {
	l.readerLock.Lock()
	defer l.readerLock.Unlock()
	l.readerCnt++
	if l.readerCnt == 1 { // 第一个读者阻止写
		l.writerLock.Lock()
	}
}

func (l *rwLock) RUnlock() {
	l.readerLock.Lock()
	defer l.readerLock.Unlock()
	l.readerCnt--
	if l.readerCnt == 0 { // 最后一个读者释放写锁
		l.writerLock.Unlock()
	}
}

func (l *rwLock) Lock() {
	l.writerLock.Lock()
}

func (l *rwLock) Unlock() {
	l.writerLock.Unlock()
}

// 模拟任务
func reader(id int, lock *rwLock, wg *sync.WaitGroup) {
	defer wg.Done()
	lock.RLock()
	fmt.Printf("Reader %d: reading...\n", id)
	time.Sleep(200 * time.Millisecond) // 模拟读耗时
	fmt.Printf("Reader %d: done\n", id)
	lock.RUnlock()
}

func writer(id int, lock *rwLock, wg *sync.WaitGroup) {
	defer wg.Done()
	lock.Lock()
	fmt.Printf("Writer %d: writing...\n", id)
	time.Sleep(300 * time.Millisecond) // 模拟写耗时
	fmt.Printf("Writer %d: done\n", id)
	lock.Unlock()
}

func main() {
	lock := NewRWLock()
	var wg sync.WaitGroup

	// 读写交替创建
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go reader(i, lock, &wg)

		wg.Add(1)
		go writer(i, lock, &wg)
	}

	wg.Wait()
}
