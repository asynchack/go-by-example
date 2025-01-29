package main

import (
	"fmt"
)

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// Because we closed the channel above, the iteration terminates after receiving the 2 elements.
	for v := range queue {
		fmt.Println(v)
	}

}
