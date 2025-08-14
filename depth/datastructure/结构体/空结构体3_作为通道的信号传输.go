package main

import "fmt"

// empty struct
var empty = struct{}{}

// Semaphore is empty type chan
type Semaphore chan struct{}

// P used to acquire n resources
func (s Semaphore) P(n int) {
	for i := 0; i < n; i++ {
		s <- empty
	}
}

// V used to release n resouces
func (s Semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

// Lock used to lock resource
func (s Semaphore) Lock() {
	s.P(1)
}

// Unlock used to unlock resource
func (s Semaphore) Unlock() {
	s.V(1)
}

// NewSemaphore return semaphore
func NewSemaphore(N int) Semaphore {
	return make(Semaphore, N)
}
func main() {
	sem := NewSemaphore(1)
	fmt.Println(1)
	sem.Lock()
	fmt.Println(2)
	sem.Lock()
	fmt.Println(3)

	sem.Unlock()

}
