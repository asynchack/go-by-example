package main

import (
	"errors"
	"fmt"
)

type argError struct {
	arg     int
	message string
}

func (a *argError) Error() string {
	return fmt.Sprintf("%s - %s", a.arg, a.message)
}

func f(arg int) (int, error) {
	if arg == 233 {
		return -1, &argError{arg: arg, message: "can`t work with.."}
	}
	return arg, nil
}

func main() {
	_, err := f(233)
	var ar *argError
	if errors.As(err, &ar) {
		fmt.Println(ar.arg)
		fmt.Println(ar.message)
	} else {
		fmt.Println("does`t match argError")
	}
}
