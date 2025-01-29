package main

import (
	"fmt"
	"time"
)

func worker(ch chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	ch <- true
}

func main() {
	ch := make(chan bool)
	go worker(ch)

	<-ch
}
