package main

import "fmt"

func vals() (int, int) {
	return 2, 10
}

func main() {
	a, b := vals()
	fmt.Println(a, b)
}
