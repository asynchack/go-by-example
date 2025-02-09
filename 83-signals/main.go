package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 声明一个有buffer的chan，并注册它可以处理的sig
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)

	done := make(chan bool, 1)

	go func() {
		<-sig
		fmt.Println()
		done <- true
	}()

	fmt.Println("waiting for sig...")
	<-done
	fmt.Println("exit.")
}
