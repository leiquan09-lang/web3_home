package main

import (
	"fmt"
	"sync"
)

const jobCount = 100

func main() {
	ch := make(chan int, 10) // 缓冲 10，生产暂时不用等消费

	var wg sync.WaitGroup
	wg.Add(2)

	// 生产者
	go func() {
		defer wg.Done()
		for i := 1; i <= jobCount; i++ {
			ch <- i
		}
		close(ch)
	}()

	// 消费者
	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println("recv:", v)
		}
	}()

	wg.Wait() // 等两边都收工
}
