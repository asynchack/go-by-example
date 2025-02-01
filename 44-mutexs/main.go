package main

import (
	"fmt"
	"sync"
)

type container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

// 多个goroutine借助锁实现并发操作的原子化
func main() {
	var c = container{
		// mu 无需init
		counters: map[string]int{"a": 0, "b": 0},
	}
	var wg sync.WaitGroup

	incCounter := func(name string, number int) {
		defer wg.Done()
		for i := 0; i < number; i++ {
			c.inc(name)
		}
	}

	wg.Add(3)
	go incCounter("a", 10000)
	go incCounter("a", 10000)
	go incCounter("b", 10000)
	wg.Wait()

	fmt.Println(c.counters)

	fmt.Println("")
}
