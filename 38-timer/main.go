package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("t1 fired")

	timer2 := time.NewTimer(3 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("t2 fired")
	}()

	tStop := timer2.Stop()
	if tStop {
		fmt.Println("cancel t2 ")
	}

	time.Sleep(time.Second * 7)

}
