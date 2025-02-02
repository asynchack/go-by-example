package main

import (
	"fmt"
)

func myPanic() {
	panic("a error")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover! err is:", r)
		}
	}()

	myPanic()

	fmt.Println("can`t be executed~")
}
