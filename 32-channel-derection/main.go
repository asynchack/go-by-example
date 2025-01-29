package main

import (
	"fmt"
)

func pingx(ping chan<- string, msg string) {
	ping <- msg
}

func pongx(ping <-chan string, pong chan<- string) {
	msg := <-ping
	pong <- msg
}

func main() {
	ping := make(chan string, 1)
	pong := make(chan string, 1)

	pingx(ping, "ping")
	pongx(ping, pong)
	fmt.Println(<-pong)
}
