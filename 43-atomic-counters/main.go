package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup

	var ops atomic.Uint64

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				ops.Add(1)
			}
		}()
	}
	wg.Wait()

	fmt.Println("counter: ", ops.Load()) // exactly 50 * 1000
}
