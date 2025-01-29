package main

import (
	"errors"
	"fmt"
)

var ErrOutOfTea = fmt.Errorf("no more tea avaliable")
var ErrPower = fmt.Errorf("can`t boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func f(arg int) (int, error) {
	if arg == 233 {
		return -1, errors.New("can`t work with 42!")
	}
	return arg + 1, nil
}
func main() {
	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("buy more tea")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("charge power")
			} else {
				fmt.Println("unknown error: %s", err)
			}
			continue
		}
		fmt.Println("tea is fine")
	}

	res, err := f(233)
	fmt.Println(res, err)
}
