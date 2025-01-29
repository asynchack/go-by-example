package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 2)

	go func() {
		time.Sleep(time.Second)
		c1 <- "done"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "done"
	}()

	for i := 0; i < 2; i++ {
		select {
		case <-c1:
			fmt.Println("c1 done")
		case <-c2:
			fmt.Println("c2 done")
		}
	}
}
