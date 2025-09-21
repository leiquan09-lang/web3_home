package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 1; i < 10; i += 2 {
			fmt.Println("奇数协程 --> ", i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 2; i < 10; i += 2 {
			fmt.Println("偶数协程 --> ", i)
		}
	}()
	wg.Wait()
}
