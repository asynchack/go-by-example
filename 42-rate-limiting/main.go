package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)
	limiter := time.Tick(200 * time.Millisecond)

	for i := range requests {
		<-limiter
		fmt.Println("handle request:", i, time.Now())
	}

	burstLimiter := make(chan time.Time, 3)

	// 预留3个buffer
	for i := 0; i < 3; i++ {
		burstLimiter <- time.Now()
	}

	burstRequests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		burstRequests <- i
	}
	close(burstRequests)

	go func() {
		for i := range time.Tick(200 * time.Millisecond) {
			burstLimiter <- i
		}
	}()

	for req := range burstRequests {
		<-burstLimiter
		fmt.Println("handle request: ", req, time.Now())
	}
}
