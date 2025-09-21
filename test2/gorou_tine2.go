package main

import (
	"fmt"
	"sync"
	"time"
)

type Task func() error

// 调度结果
type Result struct {
	Name string
	Err  error
	Cost time.Duration
}

func Scheduler(tasks map[string]Task) []Result {
	var (
		wg     sync.WaitGroup
		mu     sync.Mutex
		result []Result
	)
	wg.Add(len(tasks))

	for name, task := range tasks {
		go func(n string, t Task) {
			defer wg.Done()

			start := time.Now()
			err := t()
			cost := time.Since(start)

			mu.Lock()
			result = append(result, Result{Name: n, Err: err, Cost: cost})
			mu.Unlock()
		}(name, task)
	}

	wg.Wait()
	return result
}

// ---------- 测试 ----------
func main() {
	tasks := map[string]Task{
		"task-A": func() error { time.Sleep(100 * time.Millisecond); return nil },
		"task-B": func() error { time.Sleep(200 * time.Millisecond); return fmt.Errorf("B 失败") },
		"task-C": func() error { time.Sleep(150 * time.Millisecond); return nil },
	}

	res := Scheduler(tasks)

	for _, r := range res {
		fmt.Printf("任务 %s 耗时 %v 错误=%v\n", r.Name, r.Cost, r.Err)
	}
}
