package main

import "fmt"

func main() {
	msgs := make(chan string)
	signals := make(chan bool)

	// non-blocking read
	select {
	case msg := <-msgs:
		fmt.Println(msg)
	default:
		fmt.Println("no receive")
	}

	// non-blocking send
	select {
	case signals <- true:
		fmt.Println("send to")
	default:
		fmt.Println("no send")
	}

	// both non-blocking

	select {
	case msg := <-msgs:
		fmt.Println(msg)
	case signals <- true:
		fmt.Println("send to")
	default:
		fmt.Println("no activity")
	}
}
