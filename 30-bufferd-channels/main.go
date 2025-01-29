package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 2)

	ch <- "1"
	ch <- "2"

	res := <-ch
	res2 := <-ch
	fmt.Println(res, res2)
}
