package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "r1 done"
	}()

	select {
	case <-c1:
		fmt.Println("r1 done")
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 3")
	}

	// select {
	// case <-c1:
	// 	fmt.Println("r1 done")
	// case <-time.After(time.Second):
	// 	fmt.Println("timeout 1")
	// }
}
