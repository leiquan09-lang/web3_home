package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const (
	goroutines = 10
	loops      = 1000
)

// 共享资源
type Counter struct {
	mu sync.Mutex
	v  int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	c.v++
	c.mu.Unlock()
}
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v
}

type AtomicCounter struct {
	v int64
}

func (c *AtomicCounter) Inc() {
	atomic.AddInt64(&c.v, 1)
}
func (c *AtomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.v)
}

func main() {
	// ①  mutex 计数
	var wg sync.WaitGroup
	wg.Add(goroutines)
	mCounter := &Counter{}
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < loops; j++ {
				mCounter.Inc()
			}
		}()
	}
	wg.Wait()
	fmt.Println("Mutex  结果:", mCounter.Value()) // 10000

	// ② atomic 计数
	wg.Add(goroutines)
	aCounter := &AtomicCounter{}
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < loops; j++ {
				aCounter.Inc()
			}
		}()
	}
	wg.Wait()
	fmt.Println("Atomic 结果:", aCounter.Value()) // 10000
}
